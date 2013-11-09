package imdb

import (
	"reflect"
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
		want    Media
	}{
		{
			ID:      "rm1064868096",
			TitleID: "tt0167261",
			want: Media{
				ID:         "rm1064868096",
				TitleID:    "tt0167261",
				URL:        "http://www.imdb.com/media/rm1064868096/tt0167261",
				ContentURL: "http://ia.media-imdb.com/images/M/MV5BMTAyNDU0NjY4NTheQTJeQWpwZ15BbWU2MDk4MTY2Nw@@._V1._SX350_SY512_.jpg",
			},
		},
	} {
		got, err := NewMedia(client, tt.ID, tt.TitleID)
		if err != nil {
			t.Errorf("NewMedia(%s) error: %v", tt.ID, err)
		} else if !reflect.DeepEqual(tt.want, *got) {
			t.Errorf("NewMedia(%s) = %+v; want %+v", tt.ID, *got, tt.want)
		}
	}
}
