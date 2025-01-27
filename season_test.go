package imdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeason(t *testing.T) {
	assert := assert.New(t)
	for _, tt := range []struct {
		ID     string
		Season int
		want   Season
	}{
		{
			ID:     "tt0290988",
			Season: 5,
			want: Season{
				SeasonNumber: 5,
				ID:           "tt0290988",
				Episodes: []Episode{
					{
						SeasonNumber:  5,
						EpisodeNumber: 1,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNzMyNTlhZGUtOGQxYy00ZWNhLWJmZmEtZmEyYTU3YmYxNWI1XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732907",
						Name:          "Give Peace a Chance",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 2,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZDczMTBiNmYtZjdhMi00ZDM5LTkxODYtZDdmMjE1NzE2MmM5XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732928",
						Name:          "The Shit Puppets",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 3,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BM2YxYTI4NTYtNmJhOS00MGQ0LTg1OGEtZmU5OGI4ODlmOGQ3XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732924",
						Name:          "The Fuckin' Way She Goes",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 4,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BODAxMWYyYzQtZWU1Yi00ZjExLTg3NWYtZDMzMWZjOTU0YmRmXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732935",
						Name:          "You Got to Blame the Thing Up Here",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 5,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZWRlMGNjZDMtNjEyNi00YTNhLWE2ZDUtMTU5NTg5ZmJhNjlkXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732913",
						Name:          "Mr. Lahey Is a Fuckin' Drunk, and He Always Will Be",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 6,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BYTI1ZTcwYzMtOWNlMS00OTQ5LWJkZTYtYzhkMmRiZTA5YmM3XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732904",
						Name:          "Don't Cross the Shitline",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 7,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BYWFkNGViNmUtMTNiOS00ZDY1LThkYmItYTI5ZjBkZWNkODg2XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732929",
						Name:          "The Winds of Shit",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 8,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNmFkNTYzNDYtNDYzOC00OWVmLWE5MmItNzVmOTA3ZTk5MjJlXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732905",
						Name:          "Dressed All Over & Zesty Mordant",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 9,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BYzdhMmUwZTEtMGY2Zi00ZTQ4LTg3Y2QtMjZlYjI4ZmIxY2NkXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732908",
						Name:          "I Am the Liquor",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 10,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BOTM5ODAyNWUtMmJmNC00NTRmLTkxOTktNDQ5ZTJhOWM1NTA1XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0732927",
						Name:          "The Shit Blizzard",
					},
				},
			},
		},
		{
			ID:     "tt0290988",
			Season: 7,
			want: Season{
				SeasonNumber: 7,
				ID:           "tt0290988",
				Episodes: []Episode{
					{
						SeasonNumber:  7,
						EpisodeNumber: 1,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMWIwZWExNGUtNTJjNC00N2Q2LWFjYjItMzhhNzc2NmNlMzBhXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1000552",
						Name:          "I Fuckin' Miss Cory and Trevor",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 2,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNDJmZDMwNGYtODc5OS00ZTQ0LTk2MGQtZjNjYzEzMTNmNDk1XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1006252",
						Name:          "I Banged Lucy and Knocked Her Up... No Big Deal",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 3,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMmQyODA3OTMtMzQ0Ni00YzAyLWI0MGItYWQ4YTMxODc3OWI3XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1006254",
						Name:          "Three Good Men Are Dead",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 4,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZWFhNGJmZDgtZDFiMi00NTliLTkwNzYtMzQ5NjFmNjhmZDJiXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt0992775",
						Name:          "Friends of the Road",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 5,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMDlhZGViNTEtYjRlMS00OTYxLWE1YmUtM2E4ODFlODhhYzNjXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1006253",
						Name:          "The Mustard Tiger",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 6,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZDI4Nzc3MTMtNjM0OC00ZDJiLTgzMzctNTZjZGExOTM4MGY3XkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1026371",
						Name:          "We Can't Call People Without Wings Angels So We Call Them Friends",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 7,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZGUwMGVmNGYtZWJmMS00MTVjLTg0YjUtODMyMWRkMTEzODYzXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1026368",
						Name:          "Jump the Cheeseburger",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 8,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNGQ1MjljZTAtNmVmMS00NTAwLWJkODktYzZiN2Q0YWIyZjhhXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1026369",
						Name:          "Let the Liquor do the Thinking",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 9,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BYWY2Mzc1ZTUtODA2Ny00YTkzLTk1MGItM2ZkNWUwNTIyZTRiXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1026367",
						Name:          "Going Off the Rails on a Swayze Train",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 10,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMTE5MWVlNGYtNWViMS00NWE4LWE3NjgtMDgzMjcyMDQ0M2YxXkEyXkFqcGc@._V1_.jpg",
						ID:            "tt1026366",
						Name:          "A Shit River Runs Through It",
					},
				},
			},
		},
	} {
		got, err := NewSeason(client, tt.ID, tt.Season)
		if err != nil {
			t.Errorf("NewSeason(%v, %v) error: %v", tt.ID, tt.Season, err)
			continue
		}
		assert.Equal(tt.want, *got, "NewSeason(%v, %v) error: %v", tt.ID, tt.Season, err)
	}
}
