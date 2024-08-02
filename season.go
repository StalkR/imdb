package imdb

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	s := Season{
		SeasonNumber: season,
		ID:           id,
	}

	if err := s.Parse(c, page); err != nil {
		return nil, err
	}

	return &s, nil
}

var (
	seasonEpisodeEntryRE = regexp.MustCompile(`src="(https://m\.media-amazon\.com/images/[^"]*)" srcSet="[^"]*" sizes="[^"]*" width=".\d+"[^h]*[^r][^e][^f]f="/title/(tt.\d*)/\?ref_=ttep_ep(\d+)"[^"]*"([^"]*)"><`)
)

// Parse parses a Season from its page.
func (s *Season) Parse(c *http.Client, page []byte) error {
	episodesMatch := seasonEpisodeEntryRE.FindAllSubmatch(page, -1)

	for _, ep := range episodesMatch {
		n, _ := strconv.Atoi(string(ep[3]))
		episode := Episode{SeasonNumber: s.SeasonNumber, EpisodeNumber: n, ImageURL: string(ep[1]), ID: string(ep[2]), Name: string(ep[4])}
		s.Episodes = append(s.Episodes, episode)
	}

	return nil
}
