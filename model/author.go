package model

import (
	"errors"
)

const AuthorStr = "author"
const AuthorIdStr = "author:"

// Structure of an author
type Author struct {
	Id        string `json:"id, omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Validation of an author structure
func (author Author) Valid() error {
	if author.Firstname == "" {
		return errors.New("firstname is mandatory")
	}
	if author.Lastname == "" {
		return errors.New("lastname is mandatory")
	}
	return nil
}
