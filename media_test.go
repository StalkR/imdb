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
			ID:      "rm3523884288",
			TitleID: "tt0167261",
			want: Media{
				ID:         "rm3523884288",
				TitleID:    "tt0167261",
				URL:        "https://www.imdb.com/title/tt0167261/mediaviewer/rm3523884288",
				ContentURL: "https://m.media-amazon.com/images/M/MV5BMDY0NmI4ZjctN2VhZS00YzExLTkyZGItMTJhOTU5NTg4MDU4XkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_.jpg",
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
