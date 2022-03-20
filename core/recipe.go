package core

import (
	"errors"
)

type Recipe struct {
	ID[Recipe]
	Title       string
	Description string
	Resources   []Resource
	Tags        []Tag
}

var (
	ErrLengthLimitExceeded = errors.New("length limit exceeded")
	// ErrEmptyTitle          = errors.New("empty title is not allowed")
)
