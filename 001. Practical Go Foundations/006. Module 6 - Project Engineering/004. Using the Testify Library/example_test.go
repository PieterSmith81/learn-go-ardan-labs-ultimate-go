package nlp_test

import (
	"fmt"
	"nlp"
)

/*
Example (and test) for the Tokenize function in the nlp package.
Run "go test -v" in the terminal to see the verbose output of your Go tests.
And "go help testflag" to see all available Go test options.
*/
func ExampleTokenize() {
	tokens := nlp.Tokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who s on first]
}
