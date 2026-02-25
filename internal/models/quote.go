package models

import (
	"errors"
	"time"
)

// A quote from the database
type Quote struct {
	// Unique identifier of the Quote record.
	// I'd rather store as a string because databases that
	// support GUID IDs should use them for better concurrency.
	// But apparently Go SQL drivers don't support that? TODO: research
	Id        int64      `json:"id"`

	// Main text of the quote. At least some HTML is allowed. Sanitize before displaying.
	Body      string     `json:"body"`

	// Author of the quote. At least some HTML is allowed. Sanitize before displaying.
	Author    string     `json:"author"`

	// Moment when the quote was added to the database. Informational use only, for now.
	Created   time.Time  `json:"created"`
}


/*
	Validates the given Quote struct to ensure that required data is defined
*/
func (q *Quote) Validate() error {
	if 0 == len(q.Body) {
		return errors.New("Quotes require body text")
	}
	if 0 == len(q.Author) {
		return errors.New("Quotes require author text")
	}
	return nil
}