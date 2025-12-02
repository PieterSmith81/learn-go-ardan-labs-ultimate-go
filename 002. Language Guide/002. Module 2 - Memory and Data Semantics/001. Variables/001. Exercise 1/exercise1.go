// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var s string
	var b bool
	var i int

	// Display the value of those variables.
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(i)
	fmt.Println()

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	ss := "Hey, Ho, Let's Go!"
	bb := false
	ii := 42

	// Display the value of those variables.
	fmt.Println(ss)
	fmt.Println(bb)
	fmt.Println(ii)
	fmt.Println()

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the type and value of that variable.
	fmt.Printf("%T - %v\n", pi, pi)
}
