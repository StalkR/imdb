// Binary imdb is a simple command-line tool to demonstrate imdb package.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/StalkR/imdb"
)

func format(t *imdb.Title) string {
	name := t.Name
	if t.Year > 0 {
		name = fmt.Sprintf("%s (%d)", name, t.Year)
	}
	infos := []string{name}
	if len(t.Genres) > 0 {
		genres := t.Genres
		if len(t.Genres) > 3 {
			genres = genres[:3]
		}
		infos = append(infos, strings.Join(genres, ", "))
	}
	if len(t.Directors) > 0 {
		var directors []string
		for _, d := range t.Directors {
			directors = append(directors, d.FullName)
		}
		if len(directors) > 2 {
			directors = directors[:2]
		}
		infos = append(infos, strings.Join(directors, ", "))
	}
	if len(t.Actors) > 0 {
		var actors []string
		for _, a := range t.Actors {
			actors = append(actors, a.FullName)
		}
		if len(actors) > 3 {
			actors = actors[:3]
		}
		infos = append(infos, strings.Join(actors, ", "))
	}
	if t.Duration != "" {
		infos = append(infos, t.Duration)
	}
	if t.Rating != "" {
		infos = append(infos, t.Rating)
	}
	infos = append(infos, t.URL)
	return strings.Join(infos, " - ")
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <search title>\n", os.Args[0])
		os.Exit(1)
	}

	client := http.DefaultClient
	title, err := imdb.SearchTitle(client, flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(2)
	}
	if len(title) == 0 {
		fmt.Fprintf(os.Stderr, "Not found.")
		os.Exit(3)
	}
	t, err := imdb.NewTitle(client, title[0].ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(2)
	}

	fmt.Println(format(t))
}
