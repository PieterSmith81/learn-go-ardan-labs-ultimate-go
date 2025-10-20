package main

import "fmt"

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
