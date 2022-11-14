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
	searchExtraRE = regexp.MustCompile(`\(([^(]+)\)`)
	searchYearRE  = regexp.MustCompile(`\d{4}`)
	// Ignore roman numerals used for duplicates in a year, e.g. Title (II) (2000) (TV Series)
	searchRomanRE = regexp.MustCompile(`^[MDCLXVI]+$`)
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
		return newSearchTitle(page)
	}
	results := searchTitleRE.FindAllSubmatch(list[1], -1)
	if results == nil {
		return nil, nil // No results.
	}

	var t []Title
	for _, r := range results {
		var titleYear int
		var titleType string
		extras := searchExtraRE.FindAllSubmatch(r[3], -1)
		// expecting 0-3 max (any of roman/type/year)
		if len(extras) > 3 {
			return nil, fmt.Errorf("search: too many extras")
		}
		for i, x := range extras {
			if i == 0 && searchRomanRE.Match(x[1]) {
				continue // ignore roman numerals used for duplicates in a year
			}
			if digits := searchYearRE.FindSubmatch(x[1]); digits != nil {
				year, err := strconv.Atoi(string(digits[0]))
				if err != nil {
					return nil, err // should not happen as regexp matches digits
				}
				titleYear = year
			} else {
				titleType = string(x[1])
			}
		}
		id := string(r[1])
		t = append(t, Title{
			ID:   id,
			URL:  fmt.Sprintf(titleURL, id),
			Name: decode(string(r[2])),
			Year: titleYear,
			Type: titleType,
		})
	}
	return t, nil
}

var (
	// searchTitleListRE matches on each result in the list.
	newSearchTitleListRE = regexp.MustCompile(`<div class="ipc-metadata-list-summary-item__tc">(.*?)</div>`)
	// searchTitleRE matches on titles.
	newSearchTitleRE = regexp.MustCompile(`<a .*?href="/title/(tt\d+).*?>([^<]+)</a>`)
	newSearchYearRE  = regexp.MustCompile(`<ul\s*[^>]*>.*?<(?:span|label)\s*[^>]*>(\d{4})[^<]*</(?:span|label)>(.*?)</ul>`)
	newSearchTypeRE  = regexp.MustCompile(`<span\s*[^>]*>([^<]*)</span>`) // from within the year 2nd capture group
)

func newSearchTitle(page []byte) ([]Title, error) {
	list := newSearchTitleListRE.FindAllSubmatch(page, -1)
	if list == nil {
		return nil, NewErrParse("list")
	}

	var t []Title
	for _, e := range list {
		r := newSearchTitleRE.FindSubmatch(e[1])
		if r == nil {
			return nil, NewErrParse("list title")
		}
		id := string(r[1])

		var titleYear int
		var titleType string

		if m := newSearchYearRE.FindSubmatch(e[1]); m != nil {
			year, err := strconv.Atoi(string(m[1]))
			if err != nil {
				return nil, err // should not happen as regexp matches digits
			}
			titleYear = year
			if m := newSearchTypeRE.FindSubmatch(m[2]); m != nil {
				titleType = string(m[1])
			}
		}

		t = append(t, Title{
			ID:   id,
			URL:  fmt.Sprintf(titleURL, id),
			Name: decode(string(r[2])),
			Year: titleYear,
			Type: titleType,
		})
	}
	return t, nil
}
