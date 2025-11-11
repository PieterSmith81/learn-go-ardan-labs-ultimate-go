package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// We have 50 milliseconds to return an answer.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond) // Should work.
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond) // Should return the default bid.
	defer cancel()

	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

// If the algorithm didn't finish in time, return a default bid.
func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // Use a buffered channel with a buffer of 1 here to avoid goroutine leaks.

	go func() {
		bid := bestBid(url)
		ch <- bid
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

var defaultBid = Bid{
	AdURL: "http://adsЯus.com/default",
	Price: 3,
}

// Written by the Algo team. Time to completion varies.
type Bid struct {
	AdURL string
	Price int // In ¢
}

func bestBid(url string) Bid {
	// Simulate work.
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: "http://adsЯus.com/ad7",
		Price: 7,
	}
}
