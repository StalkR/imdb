// Binary imdb is a simple command-line tool to demonstrate imdb package.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/StalkR/imdb"
)

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

	fmt.Println(t.String())
}
