package imdb

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	for searchTerm, want := range map[string][]Title{
		"one battle after another": []Title{
			Title{
				ID:   "tt30144839",
				URL:  "https://www.imdb.com/title/tt30144839",
				Name: "One Battle After Another",
				Type: "movie",
				Year: 2025,
				Poster: Media{
					URL: "https://m.media-amazon.com/images/M/MV5BMzBkZmQ0NjMtNTZlMy00ZjdlLTg5ODUtYWFlNGM0YzE3MTg0XkEyXkFqcGc@._V1_.jpg",
				},
			},
			Title{
				ID:   "tt2191763",
				URL:  "https://www.imdb.com/title/tt2191763",
				Name: "One Thing After Another",
				Type: "short",
				Year: 1914,
			},
			Title{
				ID:   "tt13499806",
				URL:  "https://www.imdb.com/title/tt13499806",
				Name: "One Damn Thing After Another",
				Type: "short",
				Year: 2020,
			},
			Title{
				ID:   "tt15022684",
				URL:  "https://www.imdb.com/title/tt15022684",
				Name: "One After Another",
				Type: "movie",
				Year: 1961,
			},
		},
		"Letterkenny": []Title{
			{
				ID:   "tt4647692",
				URL:  "https://www.imdb.com/title/tt4647692",
				Name: "Letterkenny",
				Year: 2016,
				Type: "tvSeries",
				Poster: Media{
					ID:         "",
					TitleID:    "",
					URL:        "https://m.media-amazon.com/images/M/MV5BMWFjNzZjNGMtNmUxOC00MmZlLTljNjctOGUwNjdjZTQ5OGMyXkEyXkFqcGc@._V1_.jpg",
					ContentURL: "",
				},
			},
			{
				ID:   "tt3238856",
				URL:  "https://www.imdb.com/title/tt3238856",
				Name: "Letterkenny Problems",
				Year: 2013,
				Type: "tvSeries",
				Poster: Media{
					ID:         "",
					TitleID:    "",
					URL:        "https://m.media-amazon.com/images/M/MV5BNTQ2ZGNiOGYtYWVjYS00YTlkLTgzNDItZTIwMTdhM2FmMTc4XkEyXkFqcGc@._V1_.jpg",
					ContentURL: "",
				},
			},
			{
				ID:   "tt3913450",
				URL:  "https://www.imdb.com/title/tt3913450",
				Name: "Letterkenny People",
				Year: 2014,
				Type: "tvSeries",
				Poster: Media{
					ID:         "",
					TitleID:    "",
					URL:        "https://m.media-amazon.com/images/M/MV5BNmY5MDc1MDYtNzNjYS00M2JiLWI3ODgtMDBmMGJjYjhjYjZkXkEyXkFqcGc@._V1_.jpg",
					ContentURL: "",
				},
			},
			{
				ID:   "tt30767687",
				URL:  "https://www.imdb.com/title/tt30767687",
				Name: "The Produce Stand: Letterkenny",
				Year: 2020,
				Type: "podcastSeries",
				Poster: Media{
					ID:         "",
					TitleID:    "",
					URL:        "https://m.media-amazon.com/images/M/MV5BOGJiN2Y0ZTYtNmI0Ni00ZGRiLWFmMGYtN2I2YmUwZjM3YzdmXkEyXkFqcGc@._V1_.jpg",
					ContentURL: "",
				},
			},
			{
				ID:   "tt0038948",
				URL:  "https://www.imdb.com/title/tt0038948",
				Name: "Sister Kenny",
				Year: 1946,
				Type: "movie",
				Poster: Media{
					ID:         "",
					TitleID:    "",
					URL:        "https://m.media-amazon.com/images/M/MV5BMTFkM2E1NDAtY2M2MS00ZWFmLTg3M2EtOWYzY2Y4NzcyZjU0XkEyXkFqcGc@._V1_.jpg",
					ContentURL: "",
				},
			},
			{
				ID:   "tt5572524",
				URL:  "https://www.imdb.com/title/tt5572524",
				Name: "A Letter from Rose Kennedy",
				Year: 0,
				Type: "movie",
			},
		},
	} {
		results, err := SearchTitle(client, searchTerm)
		if err != nil {
			if err == errChallenged {
				log.Printf("test inconclusive: %v", errChallenged)
				continue
			}
			t.Errorf("SearchTitle(%v) error: %v", searchTerm, err)
			continue
		}
		if len(results) < len(want) {
			t.Errorf("SearchTitle(%v) got %v results, want at least %v", searchTerm, len(results), len(want))
			continue
		}
		for i, w := range want {
			if found := func() bool {
				for _, got := range results {
					if reflect.DeepEqual(got, w) {
						return true
					}
				}
				return false
			}(); !found {
				var resultsByLine []string
				for j, r := range results {
					resultsByLine = append(resultsByLine, fmt.Sprintf("%v: %#v", j, r))
				}
				t.Errorf("SearchTitle(%v) result #%v missing: %#v; results:\n%v", searchTerm, i, w, strings.Join(resultsByLine, "\n"))
			}
		}
	}
}
