package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/StalkR/imdb"
)

func init() {
	http.HandleFunc("/title/", handleTitle)
}

var titleRE = regexp.MustCompile(`^/title/(tt\d+)$`)

func handleTitle(w http.ResponseWriter, r *http.Request) {
	m := titleRE.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	id := m[1]
	b, err := title(id)
	if err != nil {
		log.Printf("title: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func title(id string) ([]byte, error) {
	t, err := imdb.NewTitle(http.DefaultClient, id)
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, err
}
