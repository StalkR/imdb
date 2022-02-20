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
	mediaJSONRE = regexp.MustCompile(`(?s)<script id="__NEXT_DATA__" type="application/json">(.*?)</script>`)
)

// Parse parses a Media from its page.
func (m *Media) Parse(page []byte) error {
	s := mediaJSONRE.FindSubmatch(page)
	if s == nil {
		return m.parseOld(page)
	}
	var v struct {
		Props struct {
			PageProps struct {
				ID               string `json:"id"`
				Img              string `json:"img"`
				InitialQueryData struct {
					Title struct {
						Images struct {
							Edges []struct {
								Node struct {
									Url string `json:"url"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"images"`
					} `json:"title"`
				} `json:"initialQueryData"`
			} `json:"pageProps"`
		} `json:"props"`
	}

	if err := json.Unmarshal(s[1], &v); err != nil {
		return err
	}
	m.TitleID = v.Props.PageProps.ID
	m.ID = v.Props.PageProps.Img
	m.URL = fmt.Sprintf(mediaURL, m.TitleID, m.ID)
	edges := v.Props.PageProps.InitialQueryData.Title.Images.Edges
	if len(edges) == 0 {
		return NewErrParse("content URL, empty edges")
	}
	m.ContentURL = edges[0].Node.Url
	if m.ContentURL == "" {
		return NewErrParse("content URL, empty")
	}
	return nil
}

// Regular expressions for old format, sometimes still appearing.
// Remove after a while, as well as from media_test.go.
var (
	mediaIDRE     = regexp.MustCompile(`<meta property="og:url" content="https://www.imdb.com/title/(tt\d+)/mediaviewer/(rm\d+)"`)
	mediaViewerRE = regexp.MustCompile(`(?s){'mediaviewer': (.*?)};`)
)

func (m *Media) parseOld(page []byte) error {
	// TitleID, ID, URL
	s := mediaIDRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("id")
	}
	m.TitleID = string(s[1])
	m.ID = string(s[2])
	m.URL = fmt.Sprintf(mediaURL, m.TitleID, m.ID)

	// ContentURL
	s = mediaViewerRE.FindSubmatch(page)
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
