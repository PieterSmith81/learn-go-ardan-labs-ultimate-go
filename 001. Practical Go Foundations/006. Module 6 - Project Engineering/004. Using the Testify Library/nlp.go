package nlp

import (
	"regexp"
	"strings"
)

var (
	// "Who's on first?" -> [Who s on first].
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

// Tokenize returns a slice of tokens (lower case) based on the words found in the text.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		tokens = append(tokens, token)
	}
	return tokens
}
