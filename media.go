package imdb

import (
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
	return fmt.Sprintf("IMDb %s/%s", m.ID, m.TitleID)
}

var rmRE = regexp.MustCompile(`^rm\d+$`)

const mediaURL = "http://www.imdb.com/media/%s/%s"

// NewMedia gets, parses and returns a Media by its ID and Title ID.
func NewMedia(c *http.Client, id, titleid string) (*Media, error) {
	if !rmRE.MatchString(id) || !ttRE.MatchString(titleid) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(mediaURL, id, titleid))
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
	mediaIDRE         = regexp.MustCompile(`<link rel="canonical" href="http://www.imdb.com/media/(rm\d+)/(tt\d+)"`)
	mediaContentURLRE = regexp.MustCompile(`(?s)<img .*? id="primary-img".*?src="([^"]+)`)
)

// Parse parses a Media from its page.
func (m *Media) Parse(page []byte) error {
	// ID, URL
	s := mediaIDRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("id")
	}
	m.ID = string(s[1])
	m.TitleID = string(s[2])
	m.URL = fmt.Sprintf(mediaURL, m.ID, m.TitleID)

	// ContentURL
	s = mediaContentURLRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("content URL")
	}
	if len(s[1]) == 0 {
		return NewErrParse("content URL empty")
	}
	m.ContentURL = string(s[1])

	return nil
}
