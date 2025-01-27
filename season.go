package imdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// Episode represents a single episode from a TV series.
type Episode struct {
	SeasonNumber, EpisodeNumber int
	ImageURL, ID, Name          string
}

// Season represents a season from a TV Series, with a list of episodes.
type Season struct {
	ID           string    `json:",omitempty"`
	SeasonNumber int       `json:",omitempty"`
	Episodes     []Episode `json:",omitempty"`
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
	s := Season{
		ID:           id,
		SeasonNumber: season,
	}
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

	var r struct {
		Episodes struct {
			Items []struct {
				ID          string `json:"id"`
				Type        string `json:"type"`
				Season      string `json:"season"`
				Episode     string `json:"episode"`
				TitleText   string `json:"titleText"`
				ReleaseDate struct {
					Month    int    `json:"month"`
					Day      int    `json:"day"`
					Year     int    `json:"year"`
					Typename string `json:"__typename"`
				} `json:"releaseDate"`
				ReleaseYear int `json:"releaseYear"`
				Image       struct {
					URL       string `json:"url"`
					MaxHeight int    `json:"maxHeight"`
					MaxWidth  int    `json:"maxWidth"`
					Caption   string `json:"caption"`
				} `json:"image"`
				Plot            string  `json:"plot"`
				AggregateRating float64 `json:"aggregateRating"`
				VoteCount       int     `json:"voteCount"`
				CanRate         bool    `json:"canRate"`
				ContributionURL string  `json:"ContributionUrl"`
			} `json:"items"`
			Total           int    `json:"total"`
			HasNextPage     bool   `json:"hasNextPage"`
			EndCursor       string `json:"endCursor"`
			HasRatedEpisode bool   `json:"hasRatedEpisode"`
		} `json:"episodes"`
	}

	if err := json.Unmarshal([]byte(jsonData), &r); err != nil {
		return err
	}
	for _, e := range r.Episodes.Items {
		if n, err := strconv.Atoi(e.Season); err != nil {
			return fmt.Errorf("Season(%v, %v): atoi(season %v): %v", s.ID, s.SeasonNumber, e.Season, err)
		} else if n != s.SeasonNumber {
			continue
		}
		episodeNumber, err := strconv.Atoi(e.Episode)
		if err != nil {
			return fmt.Errorf("Season(%v, %v): atoi(episode %v): %v", s.ID, s.SeasonNumber, e.Episode, err)
		}
		s.Episodes = append(s.Episodes, Episode{
			ID:            e.ID,
			SeasonNumber:  s.SeasonNumber,
			EpisodeNumber: episodeNumber,
			ImageURL:      e.Image.URL,
			Name:          e.TitleText,
		})
	}
	return nil
}
