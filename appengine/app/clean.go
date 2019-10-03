package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StalkR/aecache"
)

func init() {
	http.HandleFunc("/clean", handleClean)
}

func handleClean(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/clean" {
		http.NotFound(w, r)
		return
	}
	if r.Header.Get("X-Appengine-Cron") != "true" {
		http.Error(w, "", http.StatusForbidden)
		return
	}
	ctx := r.Context()
	if err := aecache.Clean(ctx); err != nil {
		log.Printf("clean: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "OK")
}
