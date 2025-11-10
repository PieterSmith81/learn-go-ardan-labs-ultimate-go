package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	nGR, nIter := 10, 1_000 // Number of goroutines, number of iterations.
	var wg sync.WaitGroup

	wg.Add(nGR)

	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				/* A data race condition occurs at the line below (count++).
				Several goroutines are trying to use the same resource and change it at the same time. */
				count++
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()

	/* You would expect the value printed below to be 10000, but it is not due to the race condition above.

	You can see this by running "go run -race count.go" in the terminal.
	"-race" is supported by the go run, build, and test commands.
	Rule of thumb: Always use "-race" with the go test command.	*/
	fmt.Println("Count:", count)
}
