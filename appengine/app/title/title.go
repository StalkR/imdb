// Package title implements the /title handler to query a title.
package title

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/StalkR/aecache"
	"github.com/StalkR/imdb"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/title/", handler)
}

var titleRE = regexp.MustCompile(`^/title/(tt\d+)$`)

func handler(w http.ResponseWriter, r *http.Request) {
	m := titleRE.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.Error(w, fmt.Sprintf("request does not match %s", titleRE), http.StatusNotFound)
		return
	}
	id := m[1]

	ctx := appengine.NewContext(r)
	b, err := aecache.Get(ctx, "title:"+id)
	if err != nil {
		b, err = title(ctx, id)
		if err != nil {
			log.Errorf(ctx, "title: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		go aecache.Set(ctx, "title:"+id, b, 24*time.Hour)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func title(ctx context.Context, id string) ([]byte, error) {
	ctx, _ = context.WithTimeout(ctx, 15*time.Second)
	client := &http.Client{
		Transport: &urlfetch.Transport{
			Context: ctx,
		},
	}
	t, err := imdb.NewTitle(client, id)
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, err
}
