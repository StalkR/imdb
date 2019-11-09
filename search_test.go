package imdb

import (
	"testing"
)

func TestSearchTitle(t *testing.T) {
	title := "Lord of the rings"
	r, err := SearchTitle(client, title)
	if err != nil {
		t.Fatalf("SearchTitle(%s) error: %v", title, err)
	}
	if len(r) < 50 {
		t.Fatalf("SearchTitle(%s) len < 50: %d", title, len(r))
	}
	if want := "tt0120737"; r[0].ID != want {
		t.Errorf("SearchTitle(%s)[0].ID = %s; want %s", title, r[0].ID, want)
	}
	if want := "The Lord of the Rings: The Fellowship of the Ring"; r[0].Name != want {
		t.Errorf("SearchTitle(%s)[0].Name = %s; want %s", title, r[0].Name, want)
	}
	for i, want := range []int{
		2001,
		0,
		2014,
		2017,
		2007,
	} {
		if r[i].Year != want {
			t.Errorf("SearchTitle(%s)[%d].Year = %d; want %d", title, i, r[i].Year, want)
		}
	}
	for i, want := range []string{
		"",
		"TV Series",
		"TV Series",
		"Short",
		"Video Game",
	} {
		if r[i].Type != want {
			t.Errorf("SearchTitle(%s)[%d].Type = %s; want %s", title, i, r[i].Type, want)
		}
	}
}

func TestSearchTitleUnicode(t *testing.T) {
	title := "Les Filles De L'Océan"
	r, err := SearchTitle(client, title)
	if err != nil {
		t.Fatalf("SearchTitle(%s) error: %v", title, err)
	}
	if len(r) == 0 {
		t.Fatalf("SearchTitle(%s) len = %d; want %d", title, len(r), 1)
	}
	if want := "tt0244764"; r[0].ID != want {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, want)
	}
}

func TestSearchTitlePositions(t *testing.T) {
	title := "Burlesque"
	r, err := SearchTitle(client, title)
	if err != nil {
		t.Fatalf("SearchTitle(%s) error: %v", title, err)
	}
	if len(r) < 3 {
		t.Fatalf("SearchTitle(%s) len = %d; want %d", title, len(r), 1)
	}
	if want := "tt1126591"; r[0].ID != want { // Burlesque (I) (2010)
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, want)
	}
	if want := "tt1586713"; r[1].ID != want { // Burlesque (II) (2010)
		t.Errorf("SearchTitle(%s)[1] = %s; want %s", title, r[1].ID, want)
	}
}

func TestMachete(t *testing.T) {
	title := "Machete Kills Again... In Space!"
	r, err := SearchTitle(client, title)
	if err != nil {
		t.Fatalf("SearchTitle(%s) error: %v", title, err)
	}
	if len(r) != 1 {
		t.Fatalf("SearchTitle(%s) len = %d; want %d", title, len(r), 1)
	}
	if want := "tt2002719"; r[0].ID != want {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, want)
	}
}
