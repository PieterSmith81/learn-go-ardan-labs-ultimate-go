package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic implementation of goroutines.
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := range 3 {
		go func() {
			fmt.Println("goroutine:", i)
		}()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println()

	/*
		Channels

		Channel semantics:
			- Send/receive to/from a channel will block until the opposite operation occurs.
				- Guarantee of delivery (there should always be a "receive" waiting for the "send" on the channel).
				- Or else you will receive the common "fatal error: all goroutines are asleep - deadlock!" error.
				- I.e., data has been sent onto the channel, which is a blocking operation, unless there is a "receive" to take that data off the channel.
				- Hence the need for the "send" below to be in its own goroutine (inside a self-executing anonymous function).
	*/
	ch := make(chan int)
	go func() {
		ch <- 7 // Send to a channel.
	}()
	v := <-ch // Receive from a channel.
	fmt.Println(v)
	fmt.Println()

	// Exercise - Sleep Sort.
	fmt.Println(sleepSort([]int{20, 30, 10})) // Should return [10 20 30].
}

/*
Exercise - Sleep Sort

This is a "joke" sorting algorithm - not for use in real-life situations.

- Algorithm:
  - For every value of "n" in values, spin up a goroutine that:
  - Sleeps for "n" milliseconds.
  - Sends "n" over a channel.
  - Collect all values from the channel to a slice and return it.
*/
func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, n := range values {
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n // Send to a channel.
		}()
	}

	var out []int
	for range values {
		n := <-ch // Receive from a channel.
		out = append(out, n)
	}
	return out
}
