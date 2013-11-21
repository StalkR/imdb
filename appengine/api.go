// Package appengine implements communication to imdb AppEngine app using JSON API.
package appengine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/StalkR/imdb"
)

// AppURL points to a running github.com/StalkR/imdb/appengine/app.
const AppURL = "https://movie-db-api.appspot.com"

// NewTitle obtains a Title ID with its information and returns a Title.
func NewTitle(c *http.Client, id string) (*imdb.Title, error) {
	resp, err := c.Get(fmt.Sprintf("%s/title/%s", AppURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(b))
	}
	var title imdb.Title
	if err = json.Unmarshal(b, &title); err != nil {
		return nil, err
	}
	return &title, nil
}

// SearchTitle searches a Title and returns a list of titles that matched.
func SearchTitle(c *http.Client, q string) ([]imdb.Title, error) {
	params := url.Values{}
	params.Set("q", q)
	resp, err := c.Get(fmt.Sprintf("%s/find?%s", AppURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(b))
	}
	var titles []imdb.Title
	if err = json.Unmarshal(b, &titles); err != nil {
		return nil, err
	}
	return titles, nil
}
