package imdb

import (
	"fmt"
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
	if accepted := map[string]bool{
		"tt7631058": true, // The Lord of the Rings (TV Series)
		"tt0120737": true, // The Lord of the Rings: The Fellowship of the Ring (2001)
	}; !accepted[r[0].ID] {
		t.Errorf("SearchTitle(%s)[0].ID = %v; want any of %v", title, r[0].ID, accepted)
	}
	if accepted := map[string]bool{
		"The Lord of the Rings":                             true,
		"The Lord of the Rings: The Fellowship of the Ring": true,
		"The Lord of the Rings: The Rings of Power":         true,
	}; !accepted[r[0].Name] {
		t.Errorf("SearchTitle(%s)[0].Name = %v; want any of %v", title, r[0].Name, accepted)
	}
	errors := []string{}
	for i, want := range []int{
		2022,
		2001,
		1990,
		2002,
		2017,
	} {
		if r[i].Year != want {
			errors = append(errors, fmt.Sprintf("SearchTitle(%s)[%d].Year = %d; want %d", title, i, r[i].Year, want))
		}
	}
	if len(errors) > 3 {
		t.Errorf("> 3 errors: %v", errors)
	}
	errors = []string{}
	for i, want := range []string{
		"TV Series",
		"TV Series",
		"Video Game",
		"TV Short",
		"Short",
	} {
		if r[i].Type != want {
			errors = append(errors, fmt.Sprintf("SearchTitle(%s)[%d].Type = %s; want %s", title, i, r[i].Type, want))
		}
	}
	if len(errors) > 3 {
		t.Errorf("> 3 errors: %v", errors)
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
