// Package title implements the /title handler to query a title.
package title

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"appengine"
	"appengine/urlfetch"
	"cache"
	"github.com/StalkR/imdb"
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

	c := appengine.NewContext(r)
	b, err := cache.Get(c, "title:"+id)
	if err != nil {
		b, err = title(c, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cache.Set(c, "title:"+id, b)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func title(c appengine.Context, id string) ([]byte, error) {
	client := urlfetch.Client(c)
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
