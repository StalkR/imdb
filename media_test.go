package imdb

import (
	"log"
	"testing"
)

func TestMedia(t *testing.T) {
	_, err := NewMedia(client, "wrong", "wrong")
	if err != ErrInvalidID {
		t.Errorf("NewMedia(wrong) = %v; want ErrInvalidId", err)
	}

	for _, tt := range []struct {
		ID      string
		TitleID string
		want    []Media
	}{
		{
			ID:      "rm2813508096",
			TitleID: "tt0167261",
			want: []Media{
				// new format
				Media{
					ID:         "rm2813508096",
					TitleID:    "tt0167261",
					URL:        "https://www.imdb.com/title/tt0167261/mediaviewer/rm2813508096",
					ContentURL: "https://m.media-amazon.com/images/M/MV5BMjA3NDk3NjI5MF5BMl5BanBnXkFtZTcwMTI2MjY0NA@@._V1_.jpg",
				},
				// old format
				Media{
					ID:         "rm2813508096",
					TitleID:    "tt0167261",
					URL:        "https://www.imdb.com/title/tt0167261/mediaviewer/rm2813508096",
					ContentURL: "https://m.media-amazon.com/images/M/MV5BMjA3NDk3NjI5MF5BMl5BanBnXkFtZTcwMTI2MjY0NA@@._V1_.jpg",
				},
			},
		},
	} {
		got, err := NewMedia(client, tt.ID, tt.TitleID)
		if err != nil {
			t.Errorf("NewMedia(%s) error: %v", tt.ID, err)
			continue
		}
		// test new and if failed, also old
		if err := diffStruct(*got, tt.want[0]); err != nil {
			if err2 := diffStruct(*got, tt.want[1]); err2 != nil {
				t.Errorf("NewMedia(%s): failed with new: %v - and old: %v", tt.ID, err, err2)
				continue
			}
			log.Printf("NewMedia(%s): parsing succeeded with old format", tt.ID)
		}
	}
}
