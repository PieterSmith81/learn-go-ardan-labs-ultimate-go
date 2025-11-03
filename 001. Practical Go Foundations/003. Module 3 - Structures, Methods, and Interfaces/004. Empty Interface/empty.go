package main

import "fmt"

func main() {
	var a any // any used to be the empty interface, i.e. interface{}, before Go version 1.18 when generics (and any) were added to Go.

	/* In the two examples below we are effectively bypassing the Go type system.
	So, the first RULE OF THUMB is: Don't use any (i.e., the empty interface).

	But there are two exceptions:
	- Serialization
	- Printing */
	a = 7
	fmt.Println("a:", a)

	a = "Hi"
	fmt.Println("a:", a)
	fmt.Println()

	// Get the underlying type of any interface by using type assertion.
	s := a.(string)
	fmt.Println("s:", s)

	// i := a.(int) // Will compile, but will panic at runtime (can't assert/convert a string to an int).
	// So, you need to use the comma, ok idiom to check the conversion.
	i, ok := a.(int)
	if ok {
		fmt.Println("i:", i)
	} else {
		fmt.Printf("Not an int (it's a %T).\n", a)
	}
	fmt.Println()

	// Type switch example. Similar to the if-based assertion/conversion above, but this time using a switch statement.
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("Other type - %T.\n", a)
	}
}
