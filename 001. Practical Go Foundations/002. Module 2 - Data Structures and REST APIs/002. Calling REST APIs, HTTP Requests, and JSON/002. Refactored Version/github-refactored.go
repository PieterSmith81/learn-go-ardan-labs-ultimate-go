package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	fmt.Println(UserInfo("ardanlabs"))
}

// UserInfo return the GitHub repository name and number of public repos from the GitHub API.
func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	resp, err := http.Get("https://api.github.com/users/ardanlabs")

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
