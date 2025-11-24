package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G♡!", 6)
	fmt.Println()

	// Note: The len() function in Go calculates the length of a string using UTF-8 encoding length,
	// not necessarily the number of characters you see on the screen.
	s := "G♡!"
	fmt.Println("len:", len(s))
	fmt.Println("s[1]:", s[1])
	fmt.Printf("s[1]: %c\n", s[1])
	fmt.Println()

	// Show the individual "runes" in a string and their position in the string (based on UTF-8 encoding).
	for i, c := range s {
		fmt.Printf("%c at %d\n", c, i)
	}

	/* Summary:
	- len() and s[] work with bytes (uint8).
	- ranging over a string (for, range) works with runes (int32).
	*/
}

func banner(text string, width int) {
	// BUG: len is in bytes.
	// padding := (width - len(text)) / 2 // Might not always work as expected due to UTF-8 encoding length (as per the comment above).
	padding := (width - utf8.RuneCountInString(text)) / 2

	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))
}
