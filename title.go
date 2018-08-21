package imdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// A Title represents an IMDb title (movie, series, etc.).
type Title struct {
	ID            string   `json:",omitempty"`
	URL           string   `json:",omitempty"`
	Name          string   `json:",omitempty"`
	Type          string   `json:",omitempty"`
	Year          int      `json:",omitempty"`
	Rating        string   `json:",omitempty"`
	Duration      string   `json:",omitempty"`
	Directors     []Name   `json:",omitempty"`
	Writers       []Name   `json:",omitempty"`
	Actors        []Name   `json:",omitempty"`
	Genres        []string `json:",omitempty"`
	Languages     []string `json:",omitempty"`
	Nationalities []string `json:",omitempty"`
	Description   string   `json:",omitempty"`
	Poster        Media    `json:",omitempty"`
	AKA           []string `json:",omitempty"`
}

// String formats a Title on one line.
func (t *Title) String() string {
	name := t.Name
	if t.Year > 0 {
		name = fmt.Sprintf("%s (%d)", name, t.Year)
	}
	infos := []string{name}
	if len(t.Genres) > 0 {
		genres := t.Genres
		if len(t.Genres) > 3 {
			genres = genres[:3]
		}
		infos = append(infos, strings.Join(genres, ", "))
	}
	if len(t.Directors) > 0 {
		var directors []string
		for _, d := range t.Directors {
			directors = append(directors, d.FullName)
		}
		if len(directors) > 2 {
			directors = directors[:2]
		}
		infos = append(infos, strings.Join(directors, ", "))
	}
	if len(t.Actors) > 0 {
		var actors []string
		for _, a := range t.Actors {
			actors = append(actors, a.FullName)
		}
		if len(actors) > 3 {
			actors = actors[:3]
		}
		infos = append(infos, strings.Join(actors, ", "))
	}
	if t.Duration != "" {
		infos = append(infos, t.Duration)
	}
	if t.Rating != "" {
		infos = append(infos, t.Rating)
	}
	infos = append(infos, t.URL)
	return strings.Join(infos, " - ")
}

var ttRE = regexp.MustCompile(`^tt\d+$`)

const titleURL = "https://www.imdb.com/title/%s"

// NewTitle gets, parses and returns a Title by its ID.
func NewTitle(c *http.Client, id string) (*Title, error) {
	if !ttRE.MatchString(id) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(titleURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	t := Title{}
	if err := t.Parse(page); err != nil {
		return nil, err
	}

	resp, err = c.Get(t.URL + "/releaseinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rls, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := t.ParseRls(rls); err != nil {
		return nil, err
	}

	return &t, nil
}

// Regular expressions to parse a Title.
var (
	schemaRE             = regexp.MustCompile(`(?s)<script type="application/ld\+json">(.*?)</script>`)
	titleYearRE          = regexp.MustCompile(`<a href="/year/(\d+)/`)
	titleDurationRE      = regexp.MustCompile(`<time datetime="(?:PT)?([0-9HM]+)"`)
	titleLanguagesRE     = regexp.MustCompile(`(?s)Language:</h4>(.*?)</div>`)
	titleLanguageRE      = regexp.MustCompile(`<a[^>]*>([^<]+)</a>`)
	titleNationalitiesRE = regexp.MustCompile(`href="/search/title\?country_of_origin[^"]+"[^>]+>([^<]+)`)
	titleDescriptionRE   = regexp.MustCompile(`<meta property="og:description" content="(?:(?:Created|Directed) by .*?\w\w\.\s*)*(?:With .*?\w\w\.\s*)?([^"]*)`)
	titlePosterRE        = regexp.MustCompile(`(?s)<div class="poster">\s*<a href="/title/tt\d+/mediaviewer/(rm\d+)[^"]*"[^>]*>\s*<img.*?src="([^"]+)"`)
)

type schemaJSON struct {
	Type            string `json:"@type"`
	Url             string
	Name            string
	Image           string
	Genre           genreJSON
	Actor           personsJSON
	Director        personsJSON
	Creator         personsJSON
	Description     string
	DatePublished   string
	AggregateRating struct {
		RatingValue string
	}
	Duration string
}

// Some types are either single or list, so we need to handle that.
type genreJSON []string

func (s *genreJSON) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err == nil {
		*s = []string{v}
		return nil
	}
	var w []string
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	*s = w
	return nil
}

type personsJSON []personJSON

func (s *personsJSON) UnmarshalJSON(data []byte) error {
	var v personJSON
	if err := json.Unmarshal(data, &v); err == nil {
		*s = []personJSON{v}
		return nil
	}
	var w []personJSON
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	*s = w
	return nil
}

type personJSON struct {
	Type string `json:"@type"`
	Url  string
	Name string
}

// Parse parses a Title from its page.
func (t *Title) Parse(page []byte) error {
	s := schemaRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("schema")
	}
	var v schemaJSON
	if err := json.Unmarshal(s[1], &v); err != nil {
		return NewErrParse(err.Error())
	}

	p := strings.Split(v.Url, "/")
	if len(p) != 4 {
		return NewErrParse("id")
	}
	t.ID = p[2]
	t.URL = fmt.Sprintf(titleURL, t.ID)
	t.Name = v.Name
	t.Type = v.Type

	if len(v.DatePublished) >= 4 {
		year, err := strconv.Atoi(v.DatePublished[:4])
		if err != nil {
			return NewErrParse(fmt.Sprintf("date: %v", err))
		}
		t.Year = year
	} else {
		s := titleYearRE.FindSubmatch(page)
		if s == nil {
			return NewErrParse("year")
		}
		t.Year, _ = strconv.Atoi(string(s[1])) // regexp matches digits
	}

	t.Rating = v.AggregateRating.RatingValue
	if v.Duration != "" {
		t.Duration = strings.ToLower(strings.TrimLeft(v.Duration, "PT"))
	} else {
		s := titleDurationRE.FindSubmatch(page)
		if s != nil {
			t.Duration = strings.ToLower(string(s[1]))
		}
	}

	t.Directors = nil
	for _, e := range v.Director {
		if e.Type != "Person" {
			continue
		}
		p := strings.Split(e.Url, "/")
		if len(p) != 4 {
			return NewErrParse("director id")
		}
		id := p[2]
		if nameSlice(t.Directors).Has(id) {
			continue
		}
		t.Directors = append(t.Directors, Name{
			ID:       id,
			URL:      fmt.Sprintf(nameURL, id),
			FullName: e.Name,
		})
	}

	t.Writers = nil
	for _, e := range v.Creator {
		if e.Type != "Person" {
			continue
		}
		p := strings.Split(e.Url, "/")
		if len(p) != 4 {
			return NewErrParse("writer id")
		}
		id := p[2]
		if nameSlice(t.Writers).Has(id) {
			continue
		}
		t.Writers = append(t.Writers, Name{
			ID:       id,
			URL:      fmt.Sprintf(nameURL, id),
			FullName: e.Name,
		})
	}

	t.Actors = nil
	for _, e := range v.Actor {
		if e.Type != "Person" {
			continue
		}
		p := strings.Split(e.Url, "/")
		if len(p) != 4 {
			return NewErrParse("actor id")
		}
		id := p[2]
		if nameSlice(t.Actors).Has(id) {
			continue
		}
		t.Actors = append(t.Actors, Name{
			ID:       id,
			URL:      fmt.Sprintf(nameURL, id),
			FullName: e.Name,
		})
	}

	t.Genres = v.Genre

	s = titleLanguagesRE.FindSubmatch(page)
	if s != nil {
		s := titleLanguageRE.FindAllSubmatch(s[1], -1)
		if s == nil {
			return NewErrParse("languages")
		}
		t.Languages = nil
		for _, m := range s {
			language := decode(string(m[1]))
			if stringSlice(t.Languages).Has(language) {
				continue
			}
			t.Languages = append(t.Languages, language)
		}
	}

	as := titleNationalitiesRE.FindAllSubmatch(page, -1)
	if as != nil {
		t.Nationalities = nil
		for _, m := range as {
			nationality := decode(string(m[1]))
			if stringSlice(t.Nationalities).Has(nationality) {
				continue
			}
			t.Nationalities = append(t.Nationalities, nationality)
		}
	}

	t.Description = v.Description

	s = titlePosterRE.FindSubmatch(page)
	if s != nil {
		id := string(s[1])
		t.Poster = Media{
			ID:         id,
			TitleID:    t.ID,
			URL:        fmt.Sprintf(mediaURL, t.ID, id),
			ContentURL: string(s[2]),
		}
	}

	return nil
}

// Regular expressions to parse a Title release info.
var (
	titleAKAsRE = regexp.MustCompile(`(?s)<table id="akas"(.*?)</table>`)
	titleAKARE  = regexp.MustCompile(`(?s)<td>([^<]*)</td>\s*<td>([^<]+)</td>\s*</tr>`)
)

// ParseRls parses a Title release info from its page.
func (t *Title) ParseRls(page []byte) error {
	// AKA
	b := titleAKAsRE.FindSubmatch(page)
	if b == nil {
		return nil
	}
	s := titleAKARE.FindAllSubmatch(b[1], -1)
	if s == nil {
		return NewErrParse("aka")
	}
	t.AKA = nil
	for _, m := range s {
		comment := decode(string(m[1]))
		aka := decode(string(m[2]))
		if stringSlice(t.AKA).Has(aka) {
			continue
		}
		t.AKA = append(t.AKA, aka)
	}
	sort.StringSlice(t.AKA).Sort()
	return nil
}
