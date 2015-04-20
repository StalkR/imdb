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
	id := "tt0120737"
	if r[0].ID != id {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, id)
	}
	name := "The Lord of the Rings: The Fellowship of the Ring"
	if r[0].Name != name {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].Name, name)
	}
	year := 2001
	if r[0].Year != year {
		t.Errorf("SearchTitle(%s)[0] = %d; want %d", title, r[0].Year, year)
	}
	typ := "Video Game"
	if r[1].Type != typ {
		t.Errorf("SearchTitle(%s)[1] = %s; want %s", title, r[1].Type, typ)
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
	id := "tt0244764"
	if r[0].ID != id {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, id)
	}
}

func TestSearchTitlePositions(t *testing.T) {
	title := "Burlesque"
	r, err := SearchTitle(client, title)
	if err != nil {
		t.Fatalf("SearchTitle(%s) error: %v", title, err)
	}
	if len(r) < 2 {
		t.Fatalf("SearchTitle(%s) len = %d; want %d", title, len(r), 1)
	}
	id := "tt1126591"
	if r[0].ID != id {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, id)
	}
	id = "tt1586713"
	if r[1].ID != id {
		t.Errorf("SearchTitle(%s)[1] = %s; want %s", title, r[1].ID, id)
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
	id := "tt2002719"
	if r[0].ID != id {
		t.Errorf("SearchTitle(%s)[0] = %s; want %s", title, r[0].ID, id)
	}
}
