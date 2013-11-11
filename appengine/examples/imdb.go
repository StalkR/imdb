// Binary imdb is a simple command-line tool to demonstrate imdb/appengine package.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	imdb "github.com/StalkR/imdb/appengine"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <search title>\n", os.Args[0])
		os.Exit(1)
	}

	c := http.DefaultClient
	title, err := imdb.SearchTitle(c, flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(2)
	}
	if len(title) == 0 {
		fmt.Fprintf(os.Stderr, "Not found.")
		os.Exit(3)
	}
	t, err := imdb.NewTitle(c, title[0].ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(2)
	}

	fmt.Println(t.String())
}
