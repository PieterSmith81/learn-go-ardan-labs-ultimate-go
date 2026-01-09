// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

func main() {

	// Declare an array of 5 strings set to its zero value.
	var movies [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	favMovies := [5]string{"The Matrix", "Zoolander", "Gladiator", "Anchorman", "Braveheart"}

	// Assign the populated array to the array of zero values.
	movies = favMovies

	// Iterate over the first array declared.
	// Display the string value and address of each element.

	/* Note that the for range loop's "movie" values and addresses below (the first two parameters in the Printf statement)
	use the value semantics version of the for range loop.
	So these two "movies" values and addresses are copies of each element in the original movies array.

	For more on the two versions/flavours (the value semantics and pointer semantics versions) of the for range loop,
	see the "Arrays pt.3 (Range Mechanics)" lesson in the "Ultimate Go - Language Guide" Ardan Labs course. */
	for i, movie := range movies {
		fmt.Printf("%s\t\t\t%v\t\t\t%s\t\t\t%v\n", movie, &movie, movies[i], &movies[i])
	}
}
