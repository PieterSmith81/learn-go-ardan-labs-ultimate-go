package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Note the use of buffered channels with a buffer of 1 here to avoid goroutine leaks.
	ch1, ch2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	select {
	case v := <-ch1:
		/* Possible solution to the goroutine leaks here:
		- select with timeout.
		- But using a buffered channel with a buffer of 1 is a better solution (as per the make statements at the top of this main() function). */
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
		/*
			case <-time.After(10 * time.Millisecond):
				fmt.Println("timeout")
		*/
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
