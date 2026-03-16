package imdb

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strings"
)

func decode(s string) string {
	return strings.TrimSpace(html.UnescapeString(s))
}

var stripTagsRE = regexp.MustCompile(`<[^>]*>`)

func stripTags(s string) string {
	return stripTagsRE.ReplaceAllString(s, "")
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

var (
	errForbidden  = errors.New("forbidden by AWS WAF")
	errChallenged = errors.New("challenged by AWS WAF")
)

func checkResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	if resp.StatusCode == http.StatusForbidden {
		return errForbidden
	}
	if resp.Header.Get("X-Amzn-Waf-Action") == "challenge" {
		return errChallenged
	}
	return fmt.Errorf("imdb: status not ok: %v", resp.Status)
}
