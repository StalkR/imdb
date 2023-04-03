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
	if len(r) < 10 {
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
		2003,
		2002,
		2022,
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
		"",
		"",
		"",
		"TV Series",
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
	if accepted := map[string]bool{
		"tt5761478": true, // Harlots (TV Series) (2017-2019)
		"tt0244764": true, // Rip Girls (TV Movie) (2000)
		"tt0098797": true, // Les filles de Caleb (TV Series) (1990-)
		"tt22522556": true, // Les Filles de l'Océan
	}; !accepted[r[0].ID] {
		t.Errorf("SearchTitle(%s)[0] = %v; want any of %v", title, r[0].ID, accepted)
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
	if accepted := map[string]bool{
		"tt1126591":  true, // Burlesque (I) (2010)
		"tt1586713":  true, // Burlesque (II) (2010)
		"tt11288016": true, // Jak si nepodelat zivot (2019) (TV Mini Series) aka "Burlesque"
	}; !accepted[r[0].ID] {
		t.Errorf("SearchTitle(%s)[0] = %v; want any of %v", title, r[0].ID, accepted)
	}
}

func TestMachete(t *testing.T) {
	title := "Machete Kills Again... In Space!"
	r, err := SearchTitle(client, title)
	if err != nil {
		t.Fatalf("SearchTitle(%s) error: %v", title, err)
	}
	if len(r) == 0 {
		t.Fatalf("SearchTitle(%s) len = %d; want > 0", title, len(r))
	}
	if accepted := map[string]bool{
		"tt2002719": true,
	}; !accepted[r[0].ID] {
		t.Errorf("SearchTitle(%s)[0] = %v; want any of %v", title, r[0].ID, accepted)
	}
}
