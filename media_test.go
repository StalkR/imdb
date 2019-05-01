package imdb

import "testing"

func TestMedia(t *testing.T) {
	_, err := NewMedia(client, "wrong", "wrong")
	if err != ErrInvalidID {
		t.Errorf("NewMedia(wrong) = %v; want ErrInvalidId", err)
	}

	for _, tt := range []struct {
		ID      string
		TitleID string
		want    Media
	}{
		{
			ID:      "rm2813508096",
			TitleID: "tt0167261",
			want: Media{
				ID:         "rm2813508096",
				TitleID:    "tt0167261",
				URL:        "https://www.imdb.com/title/tt0167261/mediaviewer/rm2813508096",
				ContentURL: "https://m.media-amazon.com/images/M/MV5BMjA3NDk3NjI5MF5BMl5BanBnXkFtZTcwMTI2MjY0NA@@._V1_SX1777_CR0,0,1777,760_AL_.jpg",
			},
		},
	} {
		got, err := NewMedia(client, tt.ID, tt.TitleID)
		if err != nil {
			t.Errorf("NewMedia(%s) error: %v", tt.ID, err)
		} else {
			if err := diffStruct(*got, tt.want); err != nil {
				t.Errorf("NewMedia(%s): %v", tt.ID, err)
			}
		}
	}
}
