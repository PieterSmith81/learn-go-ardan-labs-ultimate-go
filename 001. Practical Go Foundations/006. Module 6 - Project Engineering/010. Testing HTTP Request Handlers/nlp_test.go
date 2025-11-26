package nlp

import (
	"os"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// Basic test.
func TestTokenize(t *testing.T) {
	// Setup: call a function.
	// Teardown: defer or t.Cleanup.
	text := "Who's on first?"
	tokens := Tokenize(text)
	// We are now using a stemmer in the Tokenize function in nlp.go. So, we will no longer get an "s" as part of the returned tokens.
	// expected := []string{"who", "s", "on", "first"}
	expected := []string{"who", "on", "first"}
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
/*
func TestTokenizeTable(t *testing.T) {
	// Note the use of an anonymous struct below.
	var cases = []struct {
		text   string
		tokens []string
	}{
		{"Who's on first?", []string{"who", "on", "first"}},
		{"Who's on second?", []string{"who", "on", "second"}},
		{"", nil},
	}

	for _, tc := range cases {
		// Note the use of a "subtest" callback/anonymous function below.
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)

			// Before testify.
			// if !slices.Equal(tc.tokens, tokens) {
			// 	t.Fatalf("expected %#v, got %#v", tc.tokens, tokens)
			// }

			// After testify.
			require.Equal(t, tc.tokens, tokens)
		})
	}
}
*/

/*
Table driven test exercise.
Exercise: Read the test cases from tokenize_cases.toml instead of from an in-memory slice.
*/
func TestTokenizeTable(t *testing.T) {
	cases := loadTokenizeCases(t)

	for _, tc := range cases {
		name := tc.Name
		if name == "" {
			name = tc.Text
		}
		// Note the use of a "subtest" callback/anonymous function below.
		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			// TOML does not have nil, so we are handling that here.
			if tokens == nil {
				tokens = []string{} // Empty slice (i.e., a not nil slice).
			}
			// Using testify.
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

type tokCase struct {
	Text   string
	Tokens []string
	Name   string
}

func loadTokenizeCases(t *testing.T) []tokCase {
	file, err := os.Open("testdata/tokenize_cases.toml")
	// Using testify.
	require.NoError(t, err)
	defer file.Close()

	// Note the use of an anonymous struct below.
	var data struct {
		Cases []tokCase `toml:"case"`
	}
	dec := toml.NewDecoder(file)
	_, err = dec.Decode(&data)
	// Using testify.
	require.NoError(t, err)

	return data.Cases
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

/*
Fuzz testing.
Fuzz testing is great for finding edge cases that you would never have thought about.
To run a fuzz test with enough random test data for a long enough period, run the following command in the terminal:
"go test -fuzz . -fuzztime 5s -v"

Failing fuzz tests get added to the ./testdata/fuzz/[fuzz test name] subfolder.
Go fuzz testing (as per the terminal command above), will then always run the "special" fuzz test cases in these fuzz subfolders,
since it knows that these tests failed in the past (it will even run these special test cases if you are just running a regular "go test -v").
*/
/* FuzzTokenizer randomly tests that the Tokenize function (in the nlp.go file) always returns valid tokens,
i.e., exactly the correct tokens for the sentence it is "tokenizing". */
func FuzzTokenizer(f *testing.F) {
	f.Add("") // Manually test a specific test case every time this test runs (in this case, tokenizing an empty string).

	// Fuzz test.
	fn := func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			// Using testify.
			require.Contains(t, lText, tok) // Will pass the fuzz test.
			// require.Contains(t, lText, tok+"XXX") // Will simulate failing the fuzz test.
		}
	}
	f.Fuzz(fn)
}
