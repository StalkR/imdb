package imdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// A Media represents an IMDb media (poster, photos, etc.).
// It references to a Title by its ID.
type Media struct {
	ID         string `json:",omitempty"`
	TitleID    string `json:",omitempty"`
	URL        string `json:",omitempty"`
	ContentURL string `json:",omitempty"`
}

// String formats a Media.
func (m *Media) String() string {
	return fmt.Sprintf("IMDb %s/%s", m.TitleID, m.ID)
}

var rmRE = regexp.MustCompile(`^rm\d+$`)

const mediaURL = "https://www.imdb.com/title/%s/mediaviewer/%s"

// NewMedia gets, parses and returns a Media by its ID and Title ID.
func NewMedia(c *http.Client, id, titleid string) (*Media, error) {
	if !rmRE.MatchString(id) || !ttRE.MatchString(titleid) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(mediaURL, titleid, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	m := Media{}
	if err := m.Parse(page); err != nil {
		return nil, err
	}
	return &m, nil
}

// Regular expressions to identify and parse a Media.
var (
	mediaIDRE   = regexp.MustCompile(`c4:"https://www.imdb.com/title/(tt\d+)/mediaviewer/(rm\d+)"`)
	mediaJSONRE = regexp.MustCompile(`(?s){'mediaviewer': (.*?)};`)
)

// Parse parses a Media from its page.
func (m *Media) Parse(page []byte) error {
	// TitleID, ID, URL
	s := mediaIDRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("id")
	}
	m.TitleID = string(s[1])
	m.ID = string(s[2])
	m.URL = fmt.Sprintf(mediaURL, m.TitleID, m.ID)

	// ContentURL
	s = mediaJSONRE.FindSubmatch(page)
	if s == nil || len(s[1]) == 0 {
		return NewErrParse("JSON")
	}
	var r struct {
		Galleries map[string]struct {
			AllImages []struct {
				Id  string
				Src string
			}
		}
	}
	if err := json.Unmarshal(s[1], &r); err != nil {
		return NewErrParse(fmt.Sprintf("unmarshal: %v", err))
	}
	for _, e := range r.Galleries[m.TitleID].AllImages {
		if e.Id == m.ID {
			m.ContentURL = e.Src
		}
	}
	if m.ContentURL == "" {
		return NewErrParse("content URL")
	}

	return nil
}
