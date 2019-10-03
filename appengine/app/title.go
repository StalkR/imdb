package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/StalkR/aecache"
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
	ctx := r.Context()

	b, _, err := aecache.Get(ctx, "title:"+id)
	if err != nil {
		b, err = title(id)
		if err != nil {
			log.Printf("title: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := aecache.Set(ctx, "title:"+id, b, 24*time.Hour); err != nil {
			log.Printf("title: set: %v", err)
		}
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
