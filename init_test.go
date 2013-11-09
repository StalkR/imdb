package imdb

import (
	"net/http"
	"os"

	"github.com/StalkR/httpcache"
)

// client is used by tests to perform cached requests.
// If cache directory exists it is used as a persistent cache.
// Otherwise a volatile memory cache is used.
var client *http.Client

func init() {
	if _, err := os.Stat("cache"); err == nil {
		client = httpcache.NewPersistentClient("cache")
	} else {
		client = httpcache.NewVolatileClient()
	}
}
