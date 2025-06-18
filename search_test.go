package imdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	want := []Title{
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
			ID:   "tt5572524",
			URL:  "https://www.imdb.com/title/tt5572524",
			Name: "A Letter from Rose Kennedy",
			Year: 0,
			Type: "movie",
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
	}
	assert := assert.New(t)
	results, err := SearchTitle(client, "Letterkenny")
	if err != nil {
		t.Fatalf("SearchTitle(\"Letterkenny\") error: %v", err)
	}
	if len(results) > len(want) {
		results = results[:len(want)]
	}
	for i, got := range results {
		assert.Equal(want[i], got, "SearchTitle(\"Letterkenny\") result #%v error: %v", i, err)
	}
}
