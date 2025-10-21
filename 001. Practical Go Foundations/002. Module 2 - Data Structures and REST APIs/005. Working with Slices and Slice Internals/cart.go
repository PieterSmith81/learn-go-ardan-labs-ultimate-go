package main

import (
	"fmt"
	"sort"
)

// Only used in the final example ("Looping over slices revisited") at the bottom of the main() function below.
type Player struct {
	Name  string
	Score int
}

func main() {
	fmt.Println("Create a slice and assign values to it (using literal notation):")
	cart := []string{"apple", "orange", "banana"}
	fmt.Println(cart)
	fmt.Println("len:", len(cart))
	fmt.Println("cart[1]:", cart[1])
	fmt.Println()

	fmt.Println("Range over a slice's indices:")
	for i := range cart {
		fmt.Println(i)
	}
	fmt.Println()

	fmt.Println("Range over a slice's indices and values:")
	for i, c := range cart {
		fmt.Println(i, c)
	}
	fmt.Println()

	fmt.Println("Range over a slice's values:")
	for _, c := range cart {
		fmt.Println(c)
	}
	fmt.Println()

	fmt.Println("Add elements to the slice:")
	cart = append(cart, "milk")
	fmt.Println(cart)
	fmt.Println()

	fmt.Println("Slicing operator (slicing a slice):")
	/* Note the slicing operator (":")'s half open nature.
	So the first index value is included and up to, but not including, the last index value. */
	fruit := cart[:3] // Same as cart[0:3].
	fmt.Println("fruit:", fruit)
	fmt.Println()

	/* Slices are reference types, i.e. they POINT TO UNDERLYING (A.K.A. "BACKING") ARRAYS.
	IMPORTANT: So, think of slices as a "view" of an underlying array.
	Hence, in the example below, appending to the fruit slice, also appends to the cart slice
	(both the fruit and cart slices are referencing the same underlying array in memory). */
	fmt.Println("Slices are reference types:")
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit)
	fmt.Println("cart:", cart)
	fmt.Println()

	fmt.Println("Simulate how underlying arrays grow to accommodate the data in a slice:")
	var s []int
	for i := range 10_000 {
		s = appendInt(s, i)
	}
	fmt.Println(s[:10])
	fmt.Println()

	fmt.Println("Exercise - Concat slices, without using a for loop:")
	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println("concat", out) // [A B C]
	fmt.Println()

	fmt.Println("Exercise - Median:")
	values := []float64{3, 1, 2} // 2
	fmt.Println(median(values))
	values = []float64{3, 1, 2, 4} // 2.5
	fmt.Println(median(values))
	fmt.Println("values:", values)
	fmt.Println()

	fmt.Println("Looping over slices revisited:")
	players := []Player{
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// Add a bonus.
	for _, p := range players {
		p.Score += 100
	}
	/* This doesn't show a change to players' scores - Why?
	It is due to value semantics in Go for "for" loops.
	With every iteration of the for loop above, we are getting a copy of the player (stored in the p variable) that is inside the slice.
	So we are incrementing values on the copy, not on the original/source.
	Remember that Go works "by value" most of the time. */
	fmt.Println(players)

	/* Solution:
	Change the values "inside" the slice directly (kind of similar to pointer semantics). */
	for i := range players {
		players[i].Score += 100
	}
	fmt.Println(players)
}

// appendInt manually simulates slices (and their underlying arrays) growing to accommodate more data.
func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) == cap(s) {
		// No more space in underlying array. Need to reallocate and copy.
		size := 2 * (len(s) + 1)
		fmt.Println(cap(s), "->", size)
		ns := make([]int, size)
		copy(ns, s)
		s = ns[:len(s)]
	}

	s = s[:len(s)+1]
	s[i] = v
	return s
}

// concat concatenates two string type slices without using a for loop.
func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

/*
	median does the following:

- Sort the values in a slice.
- If there is odd number of values in the slice, return only the middle value from the slice.
- Else, return the average of the two middle values from the slice.
*/
func median(values []float64) float64 {
	/* BUG - The original sorting code line below will sort the "source" slice passed to this function.
	This happens since slices are reference types in Go (i.e., slices point to an underlying array in Go).
	We don't want to do this here.
	Instead, we want to create a copy of the slice, sort the copied slice, and then return the median values from the copied slice.
	I.e., we copy here in order not to mutate the input parameter. */
	// sort.Float64s(values) // This will mutate the "source" slice, but we don't want that here, so we create a copy of the "source" slice instead.
	vals := make([]float64, len(values))
	copy(vals, values)

	sort.Float64s(vals)
	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	}

	mid := (vals[i-1] + vals[i]) / 2
	return mid
}
