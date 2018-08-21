// Package find implements the /find handler to find titles.
package find

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/StalkR/aecache"
	"github.com/StalkR/imdb"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/find", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/find" {
		http.Error(w, "request does not match /find", http.StatusNotFound)
		return
	}
	values := r.URL.Query()
	query := values.Get("q")
	if query == "" {
		http.Error(w, "q is empty", http.StatusBadRequest)
		return
	}

	c := appengine.NewContext(r)
	b, err := aecache.Get(c, "find:"+query)
	if err != nil {
		b, err = find(c, query)
		if err != nil {
			c.Errorf("find: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		go aecache.Set(c, "find:"+query, b, 24*time.Hour)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func find(c context.Context, query string) ([]byte, error) {
	client := &http.Client{
		Transport: &urlfetch.Transport{
			Context:  c,
			Deadline: 15 * time.Second,
		},
	}
	titles, err := imdb.SearchTitle(client, query)
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(titles, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
