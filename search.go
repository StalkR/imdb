package imdb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const searchURL = "https://www.imdb.com/find?%s"

var (
	titleIDRE = regexp.MustCompile(`<link rel="canonical" href="https://www.imdb.com/title/(tt\d+)/"`)
	// searchTitleListRE matches on results list. (?s) for multi-line.
	searchTitleListRE = regexp.MustCompile(`(?s)<table class="findList">(.*?)</table>`)
	// searchTitleRE matches on titles.
	searchTitleRE = regexp.MustCompile(`<a href="/title/(tt\d+).*?>([^<]+)</a>([^<]+)`)
	searchYearRE  = regexp.MustCompile(`\((\d+)\)`)
	searchTypeRE  = regexp.MustCompile(`\s+\((.+)\)`)
)

// SearchTitle searches for titles matching name and returns partial Titles.
// A partial Title has only ID, URL, Name and Year set.
// A full Title can be obtained with NewTitle, at the cost of extra requests.
func SearchTitle(c *http.Client, name string) ([]Title, error) {
	// Sections: all, tt, ep, nm, co, kw, ch, vi, qu, bi, pl
	params := url.Values{"q": {name}, "s": {"tt"}}
	resp, err := c.Get(fmt.Sprintf(searchURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	list := searchTitleListRE.FindSubmatch(page)
	if list == nil {
		return nil, NewErrParse("list")
	}
	results := searchTitleRE.FindAllSubmatch(list[1], -1)
	if results == nil {
		return nil, nil // No results.
	}

	var t []Title
	for _, r := range results {
		var year int
		y := searchYearRE.FindSubmatch(r[3])
		if y != nil {
			year, _ = strconv.Atoi(string(y[1])) // Regexp matches digits.
		}
		var titleType string
		p := searchTypeRE.FindSubmatch(r[3])
		if p != nil {
			titleType = string(p[1])
		}
		id := string(r[1])
		t = append(t, Title{
			ID:   id,
			URL:  fmt.Sprintf(titleURL, id),
			Name: decode(string(r[2])),
			Year: year,
			Type: titleType,
		})
	}
	return t, nil
}
