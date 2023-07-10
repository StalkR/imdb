package imdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// A Name represents an IMDb name (actor, director, writer, etc.).
type Name struct {
	ID       string `json:",omitempty"`
	URL      string `json:",omitempty"`
	FullName string `json:",omitempty"`
}

// String formats a Name.
func (n *Name) String() string {
	return fmt.Sprintf("IMDb %s: %s", n.ID, n.FullName)
}

var nmRE = regexp.MustCompile(`^nm\d+$`)

const nameURL = "https://www.imdb.com/name/%s"

// NewName gets, parses and returns a Name by its ID.
func NewName(c *http.Client, id string) (*Name, error) {
	if !nmRE.MatchString(id) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(nameURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return nil, errors.New("forbidden (e.g. denied by AWS WAF)")
		}
		return nil, fmt.Errorf("imdb: status not ok: %v", resp.Status)
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	n := Name{}
	if err := n.Parse(page); err != nil {
		return nil, err
	}
	return &n, nil
}

// Regular expressions to parse a Name.
var (
	nameIDLinkRE = regexp.MustCompile(`<link rel="canonical" href="https://www.imdb.com/name/(nm\d+)/"`)
)

// Parse parses a Name from its page.
func (n *Name) Parse(page []byte) error {
	// ID, URL
	s := nameIDLinkRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("id")
	}
	n.ID = string(s[1])
	n.URL = fmt.Sprintf(nameURL, n.ID)

	s = schemaRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("schema")
	}
	var v schemaJSON
	if err := json.Unmarshal(s[1], &v); err != nil {
		return NewErrParse(err.Error() + "; schema was: " + string(s[1]))
	}

	n.FullName = decode(v.Name)

	return nil
}
