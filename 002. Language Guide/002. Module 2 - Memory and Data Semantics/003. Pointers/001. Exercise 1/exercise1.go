// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

// Add imports.
import "fmt"

func main() {

	// Declare an integer variable with the value of 20.
	value := 20

	// Display the address of and value of the variable.
	fmt.Println("Address of the variable:", &value, "\tValue of the variable:", value)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	p := &value // Using the short variable declaration operator.
	/* You could also have used (so, the short variable declaration/initialization above does the following):
	var p *int
	p = &value
	*/

	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Println("Address of the pointer:", &p, "\tValue of the pointer:", p, "\tValue that the pointer points to:", *p)
}
