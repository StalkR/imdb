package imdb

import (
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
	ID            string
	URL           string
	Name          string
	Type          string
	Year          int
	Rating        string
	Duration      string
	Directors     []Name
	Writers       []Name
	Actors        []Name
	Genres        []string
	Languages     []string
	Nationalities []string
	Description   string
	Poster        Media
	AKA           []string
}

// String formats a Title.
func (t *Title) String() string {
	name := t.Name
	if t.Year != 0 {
		name = fmt.Sprintf("%s (%i)", name, t.Year)
	}
	return fmt.Sprintf("IMDb %s: %s", t.ID, name)
}

var ttRE = regexp.MustCompile(`^tt\d+$`)

const titleURL = "http://www.imdb.com/title/%s"

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
	titleIDRE            = regexp.MustCompile(`<link rel="canonical" href="http://www.imdb.com/title/(tt\d+)/"`)
	titleNameRE          = regexp.MustCompile(`property=.og:title. content="(.*?) \(`)
	titleTypeRE          = regexp.MustCompile(`property=.og:title. content="[^(]+ \((.*?) [0-9]{4}.*?\)" />`)
	titleProdYearRE      = regexp.MustCompile(`property=.og:title. content="[^(]+ \(.*?([0-9]{4}).*?\)" />`)
	titlePubYearRE       = regexp.MustCompile(`itemprop="datePublished" content="([0-9]{4})`)
	titleRatingRE        = regexp.MustCompile(`star-box-giga-star">\s*([0-9.]+)`)
	titleDurationRE      = regexp.MustCompile(`itemprop="duration" datetime="(?:PT)?([0-9HM]+)"`)
	titlePersonRE        = regexp.MustCompile(`<a href="/name/(nm\d+)/.*?itemprop="name">([^<]+)`)
	titleDirectorsRE     = regexp.MustCompile(`(?s)<div class="txt-block" itemprop="director"[^>]+>(.*?)</div>`)
	titleWritersRE       = regexp.MustCompile(`(?s)<div class="txt-block" itemprop="creator"[^>]+>(.*?)</div>`)
	titleActorsRE        = regexp.MustCompile(`(?s)<td class="itemprop" itemprop="actor".*?href="/name/(nm\d+)/.*?itemprop="name">([^<]+)`)
	titleGenresRE        = regexp.MustCompile(`(?s)<div class="[^"]+" itemprop="genre">(.*?)</div>`)
	titleGenreRE         = regexp.MustCompile(`>(.*?)</a>`)
	titleLanguagesRE     = regexp.MustCompile(`(?s)Language:</h4>(.*?)</div>`)
	titleLanguageRE      = regexp.MustCompile(`itemprop=.url.>([^<]+)</a>`)
	titleNationalitiesRE = regexp.MustCompile(`href="/country/[^"]+"[^>]+>([^<]+)`)
	titleDescriptionRE   = regexp.MustCompile(`<p itemprop="description">([^<]+)`)
	titlePosterRE        = regexp.MustCompile(`(?s)id="img_primary".* href="/media/(rm\d+).*?src="([^"]+)`)
)

// Parse parses a Title from its page.
func (t *Title) Parse(page []byte) error {
	// ID, URL
	s := titleIDRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("id")
	}
	t.ID = string(s[1])
	t.URL = fmt.Sprintf(titleURL, t.ID)

	// Name
	s = titleNameRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("name")
	}
	if len(s[1]) == 0 {
		return NewErrParse("name empty")
	}
	t.Name = decode(string(s[1]))

	// Type
	s = titleTypeRE.FindSubmatch(page)
	if s != nil {
		t.Type = string(s[1])
	}

	// Year
	s = titlePubYearRE.FindSubmatch(page)
	if s == nil {
		s = titleProdYearRE.FindSubmatch(page)
	}
	if s != nil {
		t.Year, _ = strconv.Atoi(string(s[1])) // Regexp matches digits.
	}

	// Rating
	s = titleRatingRE.FindSubmatch(page)
	if s != nil {
		t.Rating = string(s[1])
	}

	// Duration
	s = titleDurationRE.FindSubmatch(page)
	if s != nil {
		t.Duration = strings.ToLower(string(s[1]))
	}

	// Directors
	s = titleDirectorsRE.FindSubmatch(page)
	if s != nil {
		s := titlePersonRE.FindAllSubmatch(s[1], -1)
		if s == nil {
			return NewErrParse("directors")
		}
		t.Directors = nil
		for _, m := range s {
			id := string(m[1])
			if nameSlice(t.Directors).Has(id) {
				continue
			}
			t.Directors = append(t.Directors, Name{
				ID:       id,
				URL:      fmt.Sprintf(nameURL, id),
				FullName: decode(string(m[2])),
			})
		}
	}

	// Writers
	s = titleWritersRE.FindSubmatch(page)
	if s != nil {
		s := titlePersonRE.FindAllSubmatch(s[1], -1)
		if s == nil {
			return NewErrParse("writers")
		}
		t.Writers = nil
		for _, m := range s {
			id := string(m[1])
			if nameSlice(t.Writers).Has(id) {
				continue
			}
			t.Writers = append(t.Writers, Name{
				ID:       id,
				URL:      fmt.Sprintf(nameURL, id),
				FullName: decode(string(m[2])),
			})
		}
	}

	// Actors
	as := titleActorsRE.FindAllSubmatch(page, -1)
	if as != nil {
		t.Actors = nil
		for _, m := range as {
			id := string(m[1])
			if nameSlice(t.Actors).Has(id) {
				continue
			}
			t.Actors = append(t.Actors, Name{
				ID:       id,
				URL:      fmt.Sprintf(nameURL, id),
				FullName: decode(string(m[2])),
			})
		}
	}

	// Genres
	s = titleGenresRE.FindSubmatch(page)
	if s != nil {
		s := titleGenreRE.FindAllSubmatch(s[1], -1)
		if s == nil {
			return NewErrParse("genres")
		}
		t.Genres = nil
		for _, m := range s {
			genre := decode(string(m[1]))
			if stringSlice(t.Genres).Has(genre) {
				continue
			}
			t.Genres = append(t.Genres, genre)
		}
	}

	// Languages
	s = titleLanguagesRE.FindSubmatch(page)
	if s != nil {
		s := titleLanguageRE.FindAllSubmatch(s[1], -1)
		if s == nil {
			return NewErrParse("languages")
		}
		t.Languages = nil
		for _, m := range s {
			genre := decode(string(m[1]))
			if stringSlice(t.Languages).Has(genre) {
				continue
			}
			t.Languages = append(t.Languages, genre)
		}
	}

	// Nationalities
	as = titleNationalitiesRE.FindAllSubmatch(page, -1)
	if as == nil {
		return NewErrParse("nationalities")
	}
	t.Nationalities = nil
	for _, m := range as {
		nationality := decode(string(m[1]))
		if stringSlice(t.Nationalities).Has(nationality) {
			continue
		}
		t.Nationalities = append(t.Nationalities, nationality)
	}

	// Description
	s = titleDescriptionRE.FindSubmatch(page)
	if s != nil {
		t.Description = decode(string(s[1]))
	}

	// Poster
	s = titlePosterRE.FindSubmatch(page)
	if s != nil {
		id := string(s[1])
		t.Poster = Media{
			ID:         id,
			TitleID:    t.ID,
			URL:        fmt.Sprintf(mediaURL, id, t.ID),
			ContentURL: string(s[2]),
		}
	}

	return nil
}

// Regular expressions to parse a Title release info.
var (
	titleAKAsRE = regexp.MustCompile(`(?s)<table id="akas"(.*?)</table>`)
	titleAKARE  = regexp.MustCompile(`(?s)<td>([^<]+)</td>\s*</tr>`)
)

// ParseRls parses a Title release info from its page.
func (t *Title) ParseRls(page []byte) error {
	// AKA
	b := titleAKAsRE.FindSubmatch(page)
	if b != nil {
		s := titleAKARE.FindAllSubmatch(b[1], -1)
		if s == nil {
			return NewErrParse("aka")
		}
		t.AKA = nil
		for _, m := range s {
			aka := decode(string(m[1]))
			if stringSlice(t.AKA).Has(aka) {
				continue
			}
			t.AKA = append(t.AKA, aka)
		}
		sort.StringSlice(t.AKA).Sort()
	}

	return nil
}
