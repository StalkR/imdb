package imdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const searchURL = "https://v3.sg.media-imdb.com/suggestion/x/%s.json?includeVideos=0"

// SearchQueryResponse represents the JSON struct returned by a query on searchURL
type SearchQueryResponse struct {
	Matches []struct {
		Image struct {
			Height   int    `json:"height"`
			ImageURL string `json:"imageUrl"`
			Width    int    `json:"width"`
		} `json:"i"`
		ID               string `json:"id"`
		Title            string `json:"l"`
		MainActors       string `json:"s"`
		Q                string `json:"q,omitempty"`
		Type             string `json:"qid,omitempty"`
		Rank             int    `json:"rank,omitempty"`
		Year             int    `json:"y,omitempty"`
		YearInProduction string `json:"yr,omitempty"`
	} `json:"d"`
	Query string `json:"q"`
	V     int    `json:"v"`
}

// SearchTitle searches for titles matching name and returns partial Titles.
// A partial Title has only ID, URL, Name and Year set.
// A full Title can be obtained with NewTitle, at the cost of extra requests.
func SearchTitle(c *http.Client, name string) ([]Title, error) {
	resp, err := c.Get(fmt.Sprintf(searchURL, name))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden (e.g. denied by AWS WAF)")
		}
		return nil, fmt.Errorf("imdb: status not ok: %v", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if decoder == nil {
		return nil, errors.New("imdb: decoder is nil")
	}

	var searchResponse SearchQueryResponse
	err = decoder.Decode(&searchResponse)
	if err != nil {
		return nil, err
	}

	var t []Title
	for _, match := range searchResponse.Matches {
		if strings.HasPrefix(match.ID, "tt") {
			t = append(t, Title{
				ID:   match.ID,
				URL:  fmt.Sprintf(titleURL, match.ID),
				Name: match.Title,
				Year: match.Year,
				Type: match.Type,
				Poster: Media{
					URL: match.Image.ImageURL,
				},
			})
		}
	}
	return t, nil
}
