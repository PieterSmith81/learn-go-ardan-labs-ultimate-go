package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Given a GitHub user login, return name and number of public repos.

func main() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s", resp.Status)
		return
	}

	ctype := resp.Header.Get("content-type")
	fmt.Println("content-type:", ctype)

	// io.Copy(os.Stdout, resp.Body) // Read and copy the entire http response body to the Stdout (i.e., display it to the screen).

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

	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(reply.Name, reply.NumRepos)
}
