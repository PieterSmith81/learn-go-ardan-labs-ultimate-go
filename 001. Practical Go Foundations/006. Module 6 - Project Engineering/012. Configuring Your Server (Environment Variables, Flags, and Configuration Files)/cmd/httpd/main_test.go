package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Test_healthHandler will always have a successful test result (even if the web server is not running).
This happens since the health function used inside the healthHandler function is fictitious (i.e., health() just returns a nil error).
*/
func Test_healthHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	api := API{log: slog.Default()}
	api.healthHandler(w, r)

	resp := w.Result()
	// Using testify.
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
