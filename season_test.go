package imdb

import "testing"

func TestSeason(t *testing.T) {
	for _, tt := range []struct {
		season int
		want   Season
	}{
		{
			season: 7,
			want: Season{
				SeasonNumber: 7,
				ID:           "tt0290988",
				Episodes: []Episode{
					{
						SeasonNumber:  7,
						EpisodeNumber: 1,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMWIwZWExNGUtNTJjNC00N2Q2LWFjYjItMzhhNzc2NmNlMzBhXkEyXkFqcGc@._V1_QL75_UY281_CR8,0,500,281_.jpg",
						ID:            "tt1000552",
						Name:          "I Fuckin&#x27; Miss Cory and Trevor",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 2,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNDJmZDMwNGYtODc5OS00ZTQ0LTk2MGQtZjNjYzEzMTNmNDk1XkEyXkFqcGc@._V1_QL75_UY281_CR8,0,500,281_.jpg",
						ID:            "tt1006252",
						Name:          "I Banged Lucy and Knocked Her Up... No Big Deal",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 3,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMmQyODA3OTMtMzQ0Ni00YzAyLWI0MGItYWQ4YTMxODc3OWI3XkEyXkFqcGc@._V1_QL75_UY281_CR10,0,500,281_.jpg",
						ID:            "tt1006254",
						Name:          "Three Good Men Are Dead",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 4,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZWFhNGJmZDgtZDFiMi00NTliLTkwNzYtMzQ5NjFmNjhmZDJiXkEyXkFqcGc@._V1_QL75_UY281_CR10,0,500,281_.jpg",
						ID:            "tt0992775",
						Name:          "Friends of the Road",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 5,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMDlhZGViNTEtYjRlMS00OTYxLWE1YmUtM2E4ODFlODhhYzNjXkEyXkFqcGc@._V1_QL75_UY281_CR8,0,500,281_.jpg",
						ID:            "tt1006253",
						Name:          "The Mustard Tiger",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 6,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZDI4Nzc3MTMtNjM0OC00ZDJiLTgzMzctNTZjZGExOTM4MGY3XkEyXkFqcGc@._V1_QL75_UY281_CR11,0,500,281_.jpg",
						ID:            "tt1026371",
						Name:          "We Can&#x27;t Call People Without Wings Angels So We Call Them Friends",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 7,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BZGUwMGVmNGYtZWJmMS00MTVjLTg0YjUtODMyMWRkMTEzODYzXkEyXkFqcGc@._V1_QL75_UY281_CR10,0,500,281_.jpg",
						ID:            "tt1026368",
						Name:          "Jump the Cheeseburger",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 8,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNGQ1MjljZTAtNmVmMS00NTAwLWJkODktYzZiN2Q0YWIyZjhhXkEyXkFqcGc@._V1_QL75_UY281_CR9,0,500,281_.jpg",
						ID:            "tt1026369",
						Name:          "Let the Liquor do the Thinking",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 9,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BYWY2Mzc1ZTUtODA2Ny00YTkzLTk1MGItM2ZkNWUwNTIyZTRiXkEyXkFqcGc@._V1_QL75_UY281_CR7,0,500,281_.jpg",
						ID:            "tt1026367",
						Name:          "Going Off the Rails on a Swayze Train",
					},
					{
						SeasonNumber:  7,
						EpisodeNumber: 10,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMTE5MWVlNGYtNWViMS00NWE4LWE3NjgtMDgzMjcyMDQ0M2YxXkEyXkFqcGc@._V1_QL75_UY281_CR11,0,500,281_.jpg",
						ID:            "tt1026366",
						Name:          "A Shit River Runs Through It",
					},
				},
			},
		},
		{
			season: 5,
			want: Season{
				SeasonNumber: 5,
				ID:           "tt0290988",
				Episodes: []Episode{
					{
						SeasonNumber:  5,
						EpisodeNumber: 1,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMTZjOGVhMTQtOGI3OC00OTczLTlmY2ItMDkwZDQzMTc1YWNiXkEyXkFqcGc@._V1_QL75_UX500_CR0,40,500,281_.jpg",
						ID:            "tt0732907",
						Name:          "Give Peace a Chance",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 2,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMDg1ZmMxZGYtODkxMS00NDZmLTgwYTYtMTVjYjY3OTMxZWFlXkEyXkFqcGc@._V1_QL75_UY281_CR7,0,500,281_.jpg",
						ID:            "tt0732928",
						Name:          "The Shit Puppets",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 3,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMmE4MDVhODctZGZmZi00MTdiLTkyZjUtYzBmNWZlYWU5ZDVlXkEyXkFqcGc@._V1_QL75_UX500_CR0,40,500,281_.jpg",
						ID:            "tt0732924",
						Name:          "The Fuckin&#x27; Way She Goes",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 4,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNmUyYWZmZTgtZGExMS00YzliLWFkZjYtZjZhNTk5YWY0NjM2XkEyXkFqcGdeQXVyMTQ2ODE0NDA@._V1_QL75_UX500_CR0,86,500,281_.jpg",
						ID:            "tt0732935",
						Name:          "You Got to Blame the Thing Up Here",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 5,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMjU3Y2IzMDMtYjFkZC00NTdhLWI1YTQtMTRhNGM2MDY2MWJhXkEyXkFqcGdeQXVyMTQ2ODE0NDA@._V1_QL75_UX500_CR0,52,500,281_.jpg",
						ID:            "tt0732913",
						Name:          "Mr. Lahey Is a Fuckin&#x27; Drunk, and He Always Will Be",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 6,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNTE1MmU3MjItZWE4ZC00NjkwLTllZDAtMjM4YTlmZDM2MDRlXkEyXkFqcGdeQXVyMTQ2ODE0NDA@._V1_QL75_UX500_CR0,46,500,281_.jpg",
						ID:            "tt0732904",
						Name:          "Don&#x27;t Cross the Shitline",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 7,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BNjZlMjg0NDAtODk5Yy00ZDc5LWFjZTctYmNlM2FhZDcxZmU3XkEyXkFqcGc@._V1_QL75_UY281_CR7,0,500,281_.jpg",
						ID:            "tt0732929",
						Name:          "The Winds of Shit",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 8,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BYTg1ZGI2YmItNDJiZC00NjJiLWE2ZTEtNTc2MGJlNzM3YjRkL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyMTQ2ODE0NDA@._V1_QL75_UX500_CR0,46,500,281_.jpg",
						ID:            "tt0732905",
						Name:          "Dressed All Over &amp; Zesty Mordant",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 9,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BMDNhZmYxMWMtYzhmYi00OGFkLThlZTgtY2EyZTA2MjU0OGQ1XkEyXkFqcGc@._V1_QL75_UY281_CR9,0,500,281_.jpg",
						ID:            "tt0732908",
						Name:          "I Am the Liquor",
					},
					{
						SeasonNumber:  5,
						EpisodeNumber: 10,
						ImageURL:      "https://m.media-amazon.com/images/M/MV5BOTM5ODAyNWUtMmJmNC00NTRmLTkxOTktNDQ5ZTJhOWM1NTA1XkEyXkFqcGc@._V1_QL75_UY281_CR8,0,500,281_.jpg",
						ID:            "tt0732927",
						Name:          "The Shit Blizzard",
					},
				},
			},
		},
	} {
		got, err := NewSeason(client, "tt0290988", tt.season)
		if err != nil {
			t.Errorf("NewSeason(\"tt0290988\", %d) error: %v", tt.season, err)
		} else {

			if err := diffStruct(*got, tt.want); err != nil {
				t.Errorf("NewSeason(\"tt0290988\", %d) error: %v", tt.season, err)
			}
		}
	}

}
