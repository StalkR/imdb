package imdb

import (
	"testing"
)

func TestSearch(t *testing.T) {
	want := []Title{
		{
			ID:     "tt4647692",
			URL:    "https://www.imdb.com/title/tt4647692",
			Name:   "Letterkenny",
			Year:   2016,
			Type:   "tvSeries",
			Poster: Media{ID: "", TitleID: "", URL: "https://m.media-amazon.com/images/M/MV5BMjM2MjE3MDktODU4Yi00YzY5LTgzMDktYWIzMTMxYmU2YzM5XkEyXkFqcGdeQXVyMTkxNjUyNQ@@._V1_.jpg", ContentURL: ""},
		},
		{
			ID:     "tt3238856",
			URL:    "https://www.imdb.com/title/tt3238856",
			Name:   "Letterkenny Problems",
			Year:   2013,
			Type:   "tvSeries",
			Poster: Media{ID: "", TitleID: "", URL: "https://m.media-amazon.com/images/M/MV5BNTQ2ZGNiOGYtYWVjYS00YTlkLTgzNDItZTIwMTdhM2FmMTc4XkEyXkFqcGc@._V1_.jpg", ContentURL: ""},
		},
		{
			ID:     "tt3913450",
			URL:    "https://www.imdb.com/title/tt3913450",
			Name:   "Letterkenny People",
			Year:   2014,
			Type:   "tvSeries",
			Poster: Media{ID: "", TitleID: "", URL: "https://m.media-amazon.com/images/M/MV5BNmY5MDc1MDYtNzNjYS00M2JiLWI3ODgtMDBmMGJjYjhjYjZkXkEyXkFqcGc@._V1_.jpg", ContentURL: ""},
		},
		{
			ID:     "tt30767687",
			URL:    "https://www.imdb.com/title/tt30767687",
			Name:   "The Produce Stand: Letterkenny",
			Year:   2020,
			Type:   "podcastSeries",
			Poster: Media{ID: "", TitleID: "", URL: "https://m.media-amazon.com/images/M/MV5BMWViNmQ4YjAtYzI4MS00YTA3LThlNzctZTkyM2RjODNjNDcxXkEyXkFqcGdeQXVyNDY3MzkyMjM@._V1_.jpg", ContentURL: ""},
		},
		{
			ID:     "tt0038948",
			URL:    "https://www.imdb.com/title/tt0038948",
			Name:   "Sister Kenny",
			Year:   1946,
			Type:   "movie",
			Poster: Media{ID: "", TitleID: "", URL: "https://m.media-amazon.com/images/M/MV5BYjZlMzc0YjgtNzgwYi00NzAyLTk4YjUtNTMyMzEwYzM2MzgwXkEyXkFqcGdeQXVyNjc0MzMzNjA@._V1_.jpg", ContentURL: ""},
		},
		{
			ID:     "tt5572524",
			URL:    "https://www.imdb.com/title/tt5572524",
			Name:   "A Letter from Rose Kennedy",
			Year:   0,
			Type:   "movie",
			Poster: Media{ID: "", TitleID: "", URL: "", ContentURL: ""},
		},
	}
	got, err := SearchTitle(client, "Letterkenny")
	if err != nil {
		t.Errorf("SearchTitle(\"Letterkenny\") error: %v", err)
	} else {
		for i, wGot := range got {
			if i >= len(want) {
				break
			}
			if err := diffStruct(wGot, want[i]); err != nil {
				t.Errorf("SearchTitle(\"Letterkenny\") result #%v error: %v", i, err)
			}
		}
	}
}
