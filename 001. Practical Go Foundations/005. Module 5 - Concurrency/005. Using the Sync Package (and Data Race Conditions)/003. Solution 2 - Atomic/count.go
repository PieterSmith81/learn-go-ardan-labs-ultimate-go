package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/* It is not really recommended using the atomic package in Go, unless you absolutely have to.
	It is a low-level package that provides atomic memory primitives that are useful for implementing synchronization algorithms.
	But it is quite complex to use and has many caveats, so should be avoided where possible.
	So, try to use a mutex instead of atomic when using the sync package/wait groups. */
	count := int64(0)       // Atomic requires you to work with concrete types.
	nGR, nIter := 10, 1_000 // Number of goroutines, number of iterations.

	var wg sync.WaitGroup

	wg.Add(nGR)

	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				// The data race conditions have now been avoided by using the atomic package.
				atomic.AddInt64(&count, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()

	// Now printing 10000 as expected, since using the atomic package fixed all the data race conditions.
	fmt.Println("Count:", count)
}
