// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := range 10 {
		numbers = append(numbers, i)
	}

	// Display each value in the slice.
	for _, number := range numbers {
		fmt.Println(number)
	}
	fmt.Println()

	// Declare a slice of strings and populate the slice with names.
	names := []string{"Neo", "Trinity", "Morpheus", "Agent Smith", "The Oracle"}

	// Display each index position and slice value.
	for i, name := range names {
		fmt.Printf("%d \t %s\n", i, name)
	}
	fmt.Println()

	// Take a slice of index 1 and 2 of the slice of strings.
	slicedNames := names[1:3]

	// Display each index position and slice values for the new slice.
	for i, name := range slicedNames {
		fmt.Printf("%d \t %s\n", i, name)
	}
}
