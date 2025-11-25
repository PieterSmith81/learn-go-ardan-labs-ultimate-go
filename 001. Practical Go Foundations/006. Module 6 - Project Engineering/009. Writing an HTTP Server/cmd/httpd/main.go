package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nlp"
	"nlp/stemmer"
	"os"
)

func main() {
	// Routing.
	// You can test the routes below using the REST Client tests in the ./requests.http file.
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /tokenize", tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", stemHandler)

	// Spin up the web server.
	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// healthHandler (GET route handler).
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		http.Error(w, "Health check failed", http.StatusInternalServerError)
		return // Always remember to return after http.Error.
	}

	fmt.Fprintf(w, "OK")
}

// tokenizeHandler (POST route handler).
func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* Usually, there are 3 steps in a route handler:
	1. Read the data, parse the data, and validate the data (DATA VALIDATION IS SUPER IMPORTANT!).
	2. Do the work.
	3. Encode the response. */

	// STEP 1:
	// Read the data.
	// TODO: Implement http.MaxBytesReader.
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read the HTTP request body", http.StatusBadRequest)
		return // Always remember to return after http.Error.
	}

	// Parse the data.
	text := string(data)

	// Validate the data.
	if len(text) == 0 {
		http.Error(w, "Empty request received", http.StatusBadRequest)
		return // Always remember to return after http.Error.
	}

	// STEP 2:
	// Do the work.
	tokens := nlp.Tokenize(text)

	// STEP 3:
	// Encode the response.
	w.Header().Set("content-type", "application/json")
	resp := map[string]any{
		"tokens": tokens,
	}
	json.NewEncoder(w).Encode(resp)

}

// stemHandler (GET dynamic route handler - can receive a dynamic path from the request URL).
func stemHandler(w http.ResponseWriter, r *http.Request) {
	word := r.PathValue("word")
	fmt.Fprintln(w, stemmer.Stem(word))
}

func health() error {
	// TODO: Implement the actual health check.
	return nil
}
