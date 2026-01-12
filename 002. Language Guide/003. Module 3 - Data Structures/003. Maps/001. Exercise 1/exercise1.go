// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
	ages := make(map[string]int)

	// Initialize some data into the map.
	ages["Pieter"] = 44
	ages["Ida"] = 1
	ages["Mabel"] = 4
	ages["Willow"] = 5
	ages["PvZ"] = 48
	/* Maps always contain unique keys.
	This is demonstrated below, where Ida's age (previously 1) is updated in the map (to 44), rather than a second "Ida" key being added to the map. */
	ages["Ida"] = 44

	// Display each key/value pair.
	/* Note that iterating over a map in Go is always done randomly...
	To see this in action, run this go file multiple times.
	You will see that the for range loop below always prints the key/value pairs from the map in a random order. */
	for k, v := range ages {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
