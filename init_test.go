package imdb

import (
	"net/http"
	"os"
	"time"

	"github.com/StalkR/httpcache"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"

const cacheTTL = 24 * time.Hour

// client is used by tests to perform cached requests.
// If cache directory exists it is used as a persistent cache.
// Otherwise a volatile memory cache is used.
var client *http.Client

func init() {
	if _, err := os.Stat("cache"); err == nil {
		client, err = httpcache.NewPersistentClient("cache", cacheTTL)
		if err != nil {
			panic(err)
		}
	} else {
		client = httpcache.NewVolatileClient(cacheTTL, 1024)
	}
	client.Transport = &customTransport{client.Transport}
}

// customTransport implements http.RoundTripper interface to add some headers.
type customTransport struct {
	http.RoundTripper
}

func (e *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Accept-Language", "en") // avoid IP-based language detection
	r.Header.Set("User-Agent", userAgent)
	return e.RoundTripper.RoundTrip(r)
}
