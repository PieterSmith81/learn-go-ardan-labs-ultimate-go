package nlp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Basic test.
func TestTokenize(t *testing.T) {
	// Setup: call a function.
	// Teardown: defer or t.Cleanup.
	text := "Who's on first?"
	tokens := Tokenize(text)
	expected := []string{"who", "s", "on", "first"}
	/*
		// Before testify.
		if !slices.Equal(expected, tokens) {
			t.Fatalf("expected %#v, got %#v", expected, tokens)
		}
	*/
	// After testify.
	require.Equal(t, expected, tokens)
}

/*
Table driven test.
Run similar tests for different test cases/scenarios/data.
*/
func TestTokenizeTable(t *testing.T) {
	// Note the use of an anonymous struct below.
	var cases = []struct {
		text   string
		tokens []string
	}{
		{"Who's on first?", []string{"who", "s", "on", "first"}},
		{"Who's on second?", []string{"who", "s", "on", "second"}},
		{"", nil},
	}

	for _, tc := range cases {
		// Note the use of a "subtest" callback/anonymous function below.
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			/*
				// Before testify.
				if !slices.Equal(tc.tokens, tokens) {
					t.Fatalf("expected %#v, got %#v", tc.tokens, tokens)
				}
			*/
			// After testify.
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

/*
Skipping tests.
Useful for having several different sets of tests.
For example, one set of tests for developers, a different set of tests for CI/CD, a different set of tests for UI testing, etc.

Ways to select tests:
- "-run regexp" flag: Using a regular expression along with the "-run" flag on the go test command.
- Build tags: Using "//go:build" comments in your test code along with the "-tags" flag on the go test command.
- Go environment variables (usually the recommended option):
	- Running different tests based on the the current machine/environment's Go environment variables (as per the example below).
	- So, is this a developer machine/environment, or a CI/CD environment, etc?
*/

/*
Skipping tests - GitHub Actions CI/CD example.
You can simulate a "Running in the CI/CD environment" test for the TestInCI test function below by running the following command in the terminal:
"CI=yes go test -v"
Whereas a normal "go test -v" will return "Not running in the CI/CD environment."
*/
var inCI = os.Getenv("CI") != "" // GitHub Actions sets this Go environment variable automatically. In Jenkins use os.GetEnv("BUILD_NUMBER").
func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("Not running in the CI/CD environment.")
	} else {
		t.Log("Running in the CI/CD environment.")
	}
}
