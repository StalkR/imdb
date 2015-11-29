// +build appengine

// Package clean implements clean up of datastore cache.
package clean

import (
	"fmt"
	"net/http"

	"github.com/StalkR/aecache"

	"appengine"
)

func init() {
	http.HandleFunc("/cache/clean", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	c := appengine.NewContext(r)
	err := aecache.GC(c)
	if err != nil {
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "success")
}
