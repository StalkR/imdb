// Package clean implements clean up of datastore cache.
package clean

import (
	"fmt"
	"net/http"

	"github.com/StalkR/aecache"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/cache/clean", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	ctx := appengine.NewContext(r)
	err := aecache.GC(ctx)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "success")
}
