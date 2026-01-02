// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import "fmt"

const (
	// Declare a constant named server of kind string and assign a value.
	server = "server001"

	// Declare a constant named port of type "integer 16" and assign a value.
	port int16 = 443
)

func main() {

	// Display the value of both server and port.
	fmt.Printf("Server (type %T):\t%v\n", server, server)
	fmt.Printf("Port (type %T):\t%v\n", port, port)
	fmt.Println()

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	result := 1 / 3.0

	// Display the value of the variable.
	fmt.Printf("Result (type %T):\t%v\n", result, result)
}
