package imdb

import (
	"errors"
	"fmt"
)

// Generic errors.
var (
	ErrInvalidID = errors.New("imdb: invalid ID")
)

// An ErrParse represents a parsing error. It implements error interface.
type ErrParse struct {
	details string
}

// NewErrParse creates a parsing error with details.
func NewErrParse(details string) ErrParse {
	return ErrParse{details: details}
}

// Error formats a parsing error.
func (e ErrParse) Error() string {
	return fmt.Sprintf("imdb: parsing error: %s", e.details)
}
