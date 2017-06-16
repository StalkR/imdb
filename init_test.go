package imdb

import (
	"net/http"
	"os"
	"time"

	"github.com/StalkR/httpcache"
)

// client is used by tests to perform cached requests.
// If cache directory exists it is used as a persistent cache.
// Otherwise a volatile memory cache is used.
var client *http.Client

var ttl = 24 * time.Hour

func init() {
	if _, err := os.Stat("cache"); err == nil {
		client, err = httpcache.NewPersistentClient("cache", ttl)
		if err != nil {
			panic(err)
		}
	} else {
		client = httpcache.NewVolatileClient(ttl, 1024)
	}
	client.Transport = &englishTransport{client.Transport}
}

// englishTransport implements http.RoundTripper interface
// to add "Accept-Language: en" header on all requests.
type englishTransport struct {
	http.RoundTripper
}

func (e *englishTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Accept-Language", "en")
	return e.RoundTripper.RoundTrip(r)
}
