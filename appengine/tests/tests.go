// Package tests implements the /tests handler to run imdb tests.
package tests

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/tests", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "TODO")
}
