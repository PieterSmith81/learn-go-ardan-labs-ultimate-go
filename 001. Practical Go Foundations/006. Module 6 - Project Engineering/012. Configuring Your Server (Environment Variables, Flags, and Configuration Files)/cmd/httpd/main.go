package main

import (
	"encoding/json"
	"expvar" // Navigate to http://localhost:8080/debug/vars to view the output of the expvar package.
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"nlp"
	"nlp/stemmer"
)

/*
Configuration.
	- Order of preference (so, "command line" has the highest preference, followed by "environment variables", etc.):
		- defaults < configuration file < environment variables < command line.

	- Configuration file:
		- Yaml, TOML, etc. (not in the Go standard library, so you'll have to use external packages to use these file types - JSON is not recommended as a configuration file format).

	- Environment variables:
		- os.Getenv

	- Command line:
		- flag package.

	- We also have external configuration packages, libraries, and frameworks like:
		- The Ardan Labs configuration package - https://pkg.go.dev/github.com/ardanlabs/conf/v3
		- Viper (for configuration).
		- Cobra (for command line).
			- Note that Viper and Cobra can work together.

	- In this project, we are just using standard Go data structures (as per the config struct below) to store our application's configuration.

	- IMPORTANT:
		- You should always validate your configuration before doing anything else.
		- You should also always do a health check on your server before doing anything else.
		- Doing these two steps first will save you a lot of trouble down the road. */

var config struct {
	Addr string
}

func main() {
	// Configuration.
	/* Command line/environment variable configuration.
	You can run the server on its default port (port 8080), by running the following command in the terminal:
	"go run ./cmd/httpd"
	You can then manually check that the server is running by navigating to: http://localhost:8080/health

	You can also start the server on a different port (for example, port 9999), by using the following command:
	"NLP_ADDR=:9999 go run ./cmd/httpd"*/
	config.Addr = os.Getenv("NLP_ADDR")
	if config.Addr == "" {
		config.Addr = ":8080"
	}

	/* Flag configuration.
	Or you can overwrite the server's starting port, as per the command line flag implemented below.
	So, for example, by running the following command in the terminal:
	"NLP_ADDR=:9999 go run ./cmd/httpd -addr :8888" */
	flag.StringVar(&config.Addr, "addr", config.Addr, "Address to listen on")
	flag.Parse()

	// TODO: Validate configuration.

	// Health check.
	if err := health(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during initial health check - %s\n", err)
		os.Exit(1)
	}

	/* Logging.
	Here we use a technique called "dependency injection" for logging.
	You can also use dependency injection for, for example, your database connection, a connection to an authentication handler, etc.
	Dependency injection allows you to pass these application-level configurations around your application. */
	api := API{
		log: slog.Default().With("app", "nlp"),
	}

	// Routing.
	// You can test the routes below using the REST Client tests in the ./requests.http file.
	http.HandleFunc("GET /health", api.healthHandler)
	http.HandleFunc("POST /tokenize", api.tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", api.stemHandler)

	// Spin up the web server.
	api.log.Info("server starting", "address", config.Addr) // Logging.
	if err := http.ListenAndServe(config.Addr, nil); err != nil {
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
