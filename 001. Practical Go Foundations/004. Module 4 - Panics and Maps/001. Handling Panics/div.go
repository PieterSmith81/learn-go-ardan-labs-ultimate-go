package main

import "fmt"

func main() {
	/*
		fmt.Println(div(7, 3)) // Works fine - integer division will return 2.
		fmt.Println(div(7, 0)) // Will panic - division by zero.
	*/

	fmt.Println(safeDiv(7, 3))
	fmt.Println(safeDiv(7, 0))
}

func safeDiv(a, b int) (q int, err error) {
	/* Catch and recover from a panic. This must happen in a deferred self-executing anonymous function.

	Also, note the use of named returns above.
	This allows us to change the the return values in the "enclosed" anonymous function below (by using Go's "lexical scoping").

	You should only use named return values in the following two cases:
	- Defer/recover - to change a returned error value (as per the example below).
	- Documentation - for "self-documenting" your return values if there are many of them. */
	defer func() {
		if e := recover(); e != nil {
			// fmt.Println("Error:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	/*
		// This is valid code, but don't do this (i.e., return using named return variable names)!
		// It makes it difficult to follow your code's logic.
		q = div(a, b)
		err = nil
		return
	*/

	// This is much clearer/logical than the above, commented out code.
	return div(a, b), nil
}

func div(a, b int) int {
	/*
		// Manual example of throwing a panic.
		// But you should never really do this.
		// Instead, you should return errors as regular values (as per the Go Proverbs).
		if b == 0 {
			panic("Division by zero.")
		}
	*/
	return a / b
}
