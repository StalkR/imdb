package imdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
)

type seasonWrapper struct {
	Season `json:"episodes"`
}

// Season represents a season from a TV Series, with a list of episodes.
type Season struct {
	Episodes        []Episode `json:"items"`
	Total           int       `json:"total"`
	HasNextPage     bool      `json:"hasNextPage"`
	EndCursor       string    `json:"endCursor"`
	HasRatedEpisode bool      `json:"hasRatedEpisode"`
}

// Episode represents a single episode from a TV series.
type Episode struct {
	ID              string      `json:"id"`
	Type            string      `json:"type"`
	Season          string      `json:"season"`
	Episode         string      `json:"episode"`
	TitleText       string      `json:"titleText"`
	ReleaseDate     ReleaseDate `json:"releaseDate"`
	ReleaseYear     int         `json:"releaseYear"`
	Image           Image       `json:"image"`
	Plot            string      `json:"plot"`
	AggregateRating float64     `json:"aggregateRating"`
	VoteCount       int         `json:"voteCount"`
	CanRate         bool        `json:"canRate"`
	ContributionUrl string      `json:"ContributionUrl"`
}

type ReleaseDate struct {
	Month    int    `json:"month"`
	Day      int    `json:"day"`
	Year     int    `json:"year"`
	Typename string `json:"__typename"`
}

type Image struct {
	URL       string `json:"url"`
	MaxHeight int    `json:"maxHeight"`
	MaxWidth  int    `json:"maxWidth"`
	Caption   string `json:"caption"`
}

const seasonURL = "https://www.imdb.com/title/%s/episodes/?season=%d"

// NewSeason gets, parses and returns a Season by its ID and season number.
func NewSeason(c *http.Client, id string, season int) (*Season, error) {
	if !ttRE.MatchString(id) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(seasonURL, id, season))
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
	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	s := Season{}

	if err := s.Parse(c, page); err != nil {
		return nil, err
	}

	return &s, nil
}

var (
	seasonInfoJSONRE = regexp.MustCompile(`],("episodes":.*}),"currentSeason`)
)

// Parse parses a Season from its page.
func (s *Season) Parse(c *http.Client, page []byte) error {
	episodesMatch := seasonInfoJSONRE.FindSubmatch(page)
	jsonData := "{" + html.UnescapeString(string(episodesMatch[1])) + "}"
	var sw seasonWrapper
	err := json.Unmarshal([]byte(jsonData), &sw)
	if err != nil {
		return err
	}
	s.Episodes = sw.Episodes
	s.HasNextPage = sw.HasNextPage
	s.HasRatedEpisode = sw.HasRatedEpisode
	s.EndCursor = sw.EndCursor
	s.Total = sw.Total

	return nil
}
