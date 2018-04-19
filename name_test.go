package imdb

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	_, err := NewName(client, "wrong")
	if err != ErrInvalidID {
		t.Errorf("NewName(wrong) = %v; want ErrInvalidId", err)
	}

	for _, tt := range []struct {
		ID   string
		want Name
	}{
		{
			ID: "nm0130952",
			want: Name{
				ID:       "nm0130952",
				URL:      "https://www.imdb.com/name/nm0130952",
				FullName: "José Calvo",
			},
		},
		{
			ID: "nm0905152",
			want: Name{
				ID:       "nm0905152",
				URL:      "https://www.imdb.com/name/nm0905152",
				FullName: "Lilly Wachowski",
			},
		},
	} {
		got, err := NewName(client, tt.ID)
		if err != nil {
			t.Errorf("NewName(%s) error: %v", tt.ID, err)
		} else if !reflect.DeepEqual(tt.want, *got) {
			t.Errorf("NewName(%s) = %+v; want %+v", tt.ID, *got, tt.want)
		}
	}
}
