package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Given a GitHub user login, return name and number of public repos.

func main() {
	/* JSON <-> GO

	Types
	string <-> string
	true/false <-> bool
	number <-> float64, float32, int, int8 ... int64, uint, uint8 ...
	array <-> []T, []any
	object <-> map[string]any, struct

	Encoding/JSON API
	JSON -> []byte -> Go: Unmarshal
	Go -> []byte -> JSON: Marshal
	JSON -> io.Reader -> Go: Decoder
	Go -> io.Writer -> JSON: Encoder
	*/

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Should work.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond) // Should time out with context deadline exceeded.
	defer cancel()
	fmt.Println(UserInfo(ctx, "ardanlabs"))
}

// UserInfo return the GitHub repository name and number of public repos from the GitHub API.
func UserInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	/* Original code commented out.
	Now replaced with deadline-enforcing, context-based request and respond versions below. */
	// resp, err := http.Get("https://api.github.com/users/ardanlabs")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)
}

// parseResponse parses/decodes the JSON response from the GitHub API.
func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(r)

	if err := dec.Decode(&reply); err != nil {
		return "", 0, err
	}

	return reply.Name, reply.NumRepos, nil
}
