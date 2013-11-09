package imdb

import (
	"html"
	"strings"
)

func decode(s string) string {
	return strings.TrimSpace(html.UnescapeString(s))
}

type nameSlice []Name

func (s nameSlice) Has(id string) bool {
	for _, v := range s {
		if v.ID == id {
			return true
		}
	}
	return false
}

type stringSlice []string

func (s stringSlice) Has(e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
