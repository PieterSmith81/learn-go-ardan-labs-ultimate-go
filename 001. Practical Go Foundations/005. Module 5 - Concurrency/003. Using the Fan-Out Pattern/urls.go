package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://ardanlabs.com",
		"https://ibm.com/no/such/page",
	}

	// Sequential version.
	start := time.Now()
	for _, url := range urls {
		stat, err := urlCheck(url)
		fmt.Printf("%q: %d (%v)\n", url, stat, err)
	}
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
	fmt.Println()

	// Concurrent, fan-out, waiting for the results, version.
	start = time.Now()
	fanOutResult(urls)
	duration = time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
	fmt.Println()

	// Concurrent, fan-out, waiting only for the goroutines to finish (i.e., not waiting for any results), version.
	start = time.Now()
	fanOutWait(urls)
	duration = time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
	fmt.Println()

	// Concurrent, fan-out, limited/pooled, waiting only for the goroutines to finish (i.e., not waiting for any results), version.
	start = time.Now()
	fanOutPool(urls)
	duration = time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
}

func urlCheck(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func fanOutResult(urls []string) {
	type result struct {
		url    string
		status int
		err    error
	}
	ch := make(chan result)

	// Here we "fan-out" the work, sending it to multiple goroutines.
	for _, url := range urls {
		go func() {
			r := result{url: url}
			defer func() { ch <- r }()
			r.status, r.err = urlCheck(url)
		}()
	}

	// Collect the results.
	for range urls {
		r := <-ch
		fmt.Printf("%q: %d (%v)\n", r.url, r.status, r.err)
	}
}

func fanOutWait(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	// Here we "fan-out" the work, sending it to multiple goroutines.
	for _, url := range urls {
		go func() {
			defer wg.Done()
			urlLog(url)
		}()
	}

	/* Wait for the goroutines to finish.
	If you need error checking in conjunction with sync.WaitGroup, then use the errgroup experimental Go package at:
	https://pkg.go.dev/golang.org/x/sync/errgroup */
	wg.Wait()
}

func fanOutPool(urls []string) {
	var wg sync.WaitGroup
	ch := make(chan string)

	// Producer
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()

	// Here we "fan-out" the work, sending it to multiple goroutines.
	// But here, we also do limiting (a.k.a. pooling) of the goroutines, only running 2 goroutines at once (as per "range size (of 2)" for loop below).
	const size = 2
	wg.Add(size)

	for range size {
		// Consumers
		go func() {
			defer wg.Done()
			for url := range ch {
				urlLog(url)
			}
		}()
	}

	/* Wait for the goroutines to finish.
	If you need error checking in conjunction with sync.WaitGroup, then use the errgroup experimental Go package at:
	https://pkg.go.dev/golang.org/x/sync/errgroup */
	wg.Wait()
}

func urlLog(url string) {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("urlLog", "url", url, "error", err)
		return
	}
	slog.Info("urlLog", "url", url, "status", resp.StatusCode)
}
