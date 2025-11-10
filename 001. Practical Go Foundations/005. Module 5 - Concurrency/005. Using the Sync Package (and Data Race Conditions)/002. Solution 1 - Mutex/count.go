package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/* Mutex (i.e., mutual exclusion).
	The Mutex type in Go (on the sync package) makes sure only one goroutine can access a variable at a time to avoid conflicts. */
	var mu sync.Mutex
	count := 0
	nGR, nIter := 10, 1_000 // Number of goroutines, number of iterations.

	var wg sync.WaitGroup

	wg.Add(nGR)

	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				// The data race conditions (on the count++ line) have now been avoided by using a mutex.
				mu.Lock()
				count++
				mu.Unlock()
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()

	// Now printing 10000 as expected, since using a mutex fixed all the data race conditions.
	fmt.Println("Count:", count)
}
