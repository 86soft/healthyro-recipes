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

func (r *Recipe) UpdateTitle(title string) error {
	if len(title) > TitleLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.Title = title
	return nil
}

func (r *Recipe) UpdateDescription(description string) error {
	if len(description) > DescriptionLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.Description = description
	return nil
}
