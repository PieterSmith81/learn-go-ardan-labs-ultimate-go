package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic implementation of goroutines..
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
			- Send/receive to/from a channel will block until the opposite operation occurs(*).
				- Guarantee of delivery (there should always be a "receive" waiting for the "send" on the channel).
				- Or else you will receive the common "fatal error: all goroutines are asleep - deadlock!" error.
				- I.e., data has been sent onto the channel, which is a blocking operation, unless there is a "receive" to take that data off the channel.
				- Hence the need for the "send" below to be in its own goroutine (inside a self-executing anonymous function).
				- * A buffered channel has "n" non-blocking sends (this affects only sends, not receives, and you lose the guarantee of delivery).
			- Receiving from a closed channel will return the zero value for the channel (of whatever type that channel is storing) without blocking (or panicking).
				- You can use the comma, ok idiom to see if a channel is closed or not when it has a zero value.
			- Sending to a closed channel will panic.
			- Closing a closed or nil channel will panic.
			- Send/receive to a nil channel will block forever.
			- You don't need to explicitly close channels if your code doesn't require it (as per the "channel ownership" concept in Go).
				- I.e., "open" channels, like in the "Basic implementation of channels" and "Sleep Sort" examples below, do not need to be closed explicitly.
				- So, not closing "open" channels, will not cause memory leaks or any other issues.
	*/
	// Basic implementation of channels.
	ch := make(chan int)
	go func() {
		ch <- 7 // Send to a channel.
	}()
	v := <-ch // Receive from a channel.
	fmt.Println(v)
	fmt.Println()

	// Exercise - Sleep Sort.
	fmt.Println(sleepSort([]int{20, 30, 10})) // Should return [10 20 30].
	fmt.Println()

	// Iterating over, and closing channels.
	go func() {
		for i := range 4 {
			ch <- i
		}
		close(ch) // Closing a channel is a "signal" that nothing else is coming onto that channel.
	}()

	// For, ranging over a channel automatically "receives" values "off" that channel.
	for v := range ch {
		fmt.Println(">>", v)
	}
	fmt.Println()

	// Closed channels.
	// Receive from a closed channel will return the zero value for the channel (of what ever type that channel is storing) without blocking.
	v = <-ch                  // The channel is closed at this stage.
	fmt.Println("Closed:", v) // Does not panic (as per the reason above).

	// You can use the comma, ok idiom to see if a channel is closed or not when it has a zero value.
	v, ok := <-ch                           // The channel is still closed at this stage.
	fmt.Println("Closed:", v, " - ok:", ok) // The false ok value here indicates that the channel has been closed (and is no longer "ok").

	// Nil channels.
	// var ch chan int // ch is nil.
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
