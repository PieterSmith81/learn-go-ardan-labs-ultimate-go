package main

// Task: Find the N most common words in the file sherlock.txt.

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

/*
	Use a regular expression to find individual words in the text.

There are probably better ways to do this than using regular expressions (like, pre-written Go packages that specialize in word detection).
But for the purpose of this exercise we are just going to use a regular expression.
*/
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
	// mapDemo()

	// Open a file.
	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Read from the file, line by line, using a map to count the number of distinct words as you go.
	freq := make(map[string]int) // word -> count

	s := bufio.NewScanner(file)
	for s.Scan() {
		// Extract the individual words from each line.
		words := wordRe.FindAllString(s.Text(), -1)

		for _, word := range words {
			// Increment the word's count inside the map.
			freq[strings.ToLower(word)]++
		}
	}

	if err := s.Err(); err != nil {
		fmt.Println("Error reading from the file:", err)
		return
	}

	// Print the top N most common words in the file.
	top := topN(freq, 10)
	fmt.Println(top)
}

// topN returns the "n" most common words from freq.
func topN(freq map[string]int, n int) []string {
	words := slices.Collect(maps.Keys(freq))
	sort.Slice(words, func(i, j int) bool {
		wi, wj := words[i], words[j]
		// Sort in reverse order.
		return freq[wi] > freq[wj]
	})

	n = min(n, len(words))
	return words[:n]
}

// mapDemo demonstrated some of the basic operations that can be performed on maps in Go.
func mapDemo() {
	heros := map[string]string{
		// hero -> name
		"Superman":     "Clark",
		"Wonder Woman": "Diana",
		"Batman":       "Bruce",
	}

	// Range over only the keys.
	for k := range heros {
		fmt.Println(k)
	}
	fmt.Println()

	// Range over only the values.
	for _, v := range heros {
		fmt.Println(v)
	}
	fmt.Println()

	// Range over the keys and values.
	for k, v := range heros {
		fmt.Println(v, "is", k)
	}
	fmt.Println()

	// Access the value for a specific key.
	n := heros["Batman"]
	fmt.Println(n)

	// Access the value of a key that doesn't exist in a map (returns the zero value for that key's type).
	n = heros["Aquaman"] // Returns an empty string ("").
	fmt.Printf("%q\n", n)

	// Use the comma, ok idiom to find if a key exists in a map.
	n, ok := heros["Aquaman"]
	if ok {
		fmt.Printf("%q\n", n)
	} else {
		fmt.Println("Aquaman not found.")
	}
	fmt.Println()

	// Remove a key/value pair from a map.
	delete(heros, "Batman")
	fmt.Println(heros)
}
