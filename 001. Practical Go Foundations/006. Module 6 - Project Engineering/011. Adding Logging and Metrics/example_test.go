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

	/* The desired output below used to be "[who s on first]" in previous versions of this project.
	But now we are using a stemmer (see the ./stemmer package and related code in /nlp.go).
	So, now we expect an output of "[who on first]". */

	// Output:
	// [who on first]

}
