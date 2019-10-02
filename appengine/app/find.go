package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/StalkR/imdb"
)

func init() {
	http.HandleFunc("/find", handleFind)
}

func handleFind(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/find" {
		http.NotFound(w, r)
		return
	}
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing ?q=", http.StatusBadRequest)
		return
	}
	b, err := find(query)
	if err != nil {
		log.Printf("find: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func find(query string) ([]byte, error) {
	titles, err := imdb.SearchTitle(http.DefaultClient, query)
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(titles, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
