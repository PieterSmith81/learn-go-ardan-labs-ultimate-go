package main

import (
	"encoding/json"
	"expvar" // Navigate to http://localhost:8080/debug/vars to view the output of the expvar package.
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"nlp"
	"nlp/stemmer"
)

func main() {
	api := API{
		log: slog.Default().With("app", "nlp"),
	}

	// Routing.
	// You can test the routes below using the REST Client tests in the ./requests.http file.
	http.HandleFunc("GET /health", api.healthHandler)
	http.HandleFunc("POST /tokenize", api.tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", api.stemHandler)

	// Spin up the web server.
	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// healthHandler (GET route handler).
func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		a.log.Error("health", "error", err) // Logging.
		http.Error(w, "Health check failed", http.StatusInternalServerError)
		return // Always remember to return after http.Error.
	}

	fmt.Fprintf(w, "OK")
}

// tokenizeHandler (POST route handler).
func (a *API) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* Usually, there are 3 steps in a route handler:
	1. Read the data, parse the data, and validate the data (DATA VALIDATION IS SUPER IMPORTANT!).
	2. Do the work.
	3. Encode the response. */

	// STEP 1:
	// Read the data.
	// TODO: Implement http.MaxBytesReader.
	data, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error("read", "error", err, "remote", r.RemoteAddr) // Logging.
		http.Error(w, "Can't read the HTTP request body", http.StatusBadRequest)
		return // Always remember to return after http.Error.
	}

	// Parse the data.
	text := string(data)

	// Validate the data.
	if len(text) == 0 {
		a.log.Error("read", "error", "empty request") // Logging.
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
func (a *API) stemHandler(w http.ResponseWriter, r *http.Request) {
	/* Metrics.
	Navigate to http://localhost:8080/debug/vars to view the output of the expvar package,
	including the custom metric below. */
	stemCalls.Add(1) // Metrics.
	word := r.PathValue("word")
	a.log.Info("stem", "word", word) // Logging.
	fmt.Fprintln(w, stemmer.Stem(word))
}

// Helper functions.
func health() error {
	// TODO: Implement the actual health check.
	return nil
}

// Logging.
type API struct {
	log *slog.Logger
}

// Metrics.
var (
	stemCalls = expvar.NewInt("stem.calls")
)
