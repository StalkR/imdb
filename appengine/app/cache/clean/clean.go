// Package clean implements clean up of datastore cache.
package clean

import (
	"fmt"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"cache"
)

func init() {
	http.HandleFunc("/cache/clean", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	c := appengine.NewContext(r)

	var values []cache.Item
	keys, err := datastore.NewQuery("Item").GetAll(c, &values)
	if err != nil {
		c.Errorf("clean: : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Got %d item(s)\n", len(keys))

	var clean []*datastore.Key
	for i := 0; i < len(keys); i++ {
		e := values[i]
		if !e.Expires.IsZero() && e.Expires.Before(time.Now()) {
			clean = append(clean, keys[i])
		}
	}
	if len(clean) == 0 {
		fmt.Fprintf(w, "Nothing to clean\n")
		return
	}
	fmt.Fprintf(w, "%d to delete\n", len(clean))

	err = datastore.DeleteMulti(c, clean)
	if err != nil {
		c.Errorf("clean: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Deleted %d item(s)\n", len(clean))
	c.Infof("Deleted %d item(s)", len(clean))
}
