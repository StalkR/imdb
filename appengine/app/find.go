package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/StalkR/aecache"
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
	ctx := r.Context()

	b, _, err := aecache.Get(ctx, "find:"+query)
	if err != nil {
		b, err = find(query)
		if err != nil {
			log.Printf("find: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := aecache.Set(ctx, "find:"+query, b, 24*time.Hour); err != nil {
			log.Printf("find: set: %v", err)
		}
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
