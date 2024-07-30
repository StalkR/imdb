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

type Episode struct {
	SeasonNumber, EpisodeNumber int
	Title                       Title
}

type Season struct {
	ID           string    `json:",omitempty"`
	SeasonNumber int       `json:",omitempty"`
	Episodes     []Episode `json:",omitempty"`
}

const SeasonURL = "https://www.imdb.com/title/%s/episodes/?season=%d"

func NewSeason(c *http.Client, id string, season int) (*Season, error) {
	if !ttRE.MatchString(id) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(SeasonURL, id, season))
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
	seasonEpisodeEntryRE = regexp.MustCompile(`<a class="[^"]*" href="/title/(tt.\d*)/\?ref_=ttep_ep(\d+)" aria-label="(.*?)">`)
)

func (s *Season) Parse(c *http.Client, page []byte) error {
	episodesMatch := seasonEpisodeEntryRE.FindAllSubmatch(page, -1)

	for _, ep := range episodesMatch {
		n, _ := strconv.Atoi(string(ep[2]))
		tit, err := NewTitle(c, string(ep[1]))
		if err != nil {
			return err
		}
		episode := Episode{SeasonNumber: s.SeasonNumber, EpisodeNumber: n, Title: *tit}
		s.Episodes = append(s.Episodes, episode)
	}

	return nil
}
