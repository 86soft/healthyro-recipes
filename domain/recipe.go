package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Recipe struct {
	RID
	title        string
	description  string
	externalLink string
}

// NilRecipe is only for returning values on error
var NilRecipe = Recipe{}

var (
	ErrLengthLimitExceeded = errors.New("length limit exceeded")
	ErrEmptyTitle          = errors.New("empty title is not allowed")
)

func NewRecipe(title string, description string, externalLink string) (Recipe, error) {
	if title == "" {
		return NilRecipe, ErrEmptyTitle
	}
	return Recipe{
		title:        title,
		description:  description,
		externalLink: externalLink,
	}, nil
}

// UnmarshalRecipe is used only for unmarshalling Recipe from db
func UnmarshalRecipe(id RID, title string, description string, externalLink string) Recipe {
	return Recipe{
		RID:          id,
		title:        title,
		description:  description,
		externalLink: externalLink,
	}
}

func (r *Recipe) UpdateTitle(title string) error {
	if len(title) > titleLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.title = title
	return nil
}

func (r *Recipe) UpdateDescription(description string) error {
	if len(description) > descriptionLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.description = description
	return nil
}

func (r *Recipe) UpdateExternalLink(externalLink string) error {
	if len(externalLink) > externalLinkLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.externalLink = externalLink
	return nil
}

func (r *Recipe) RecipeID() uuid.UUID {
	return r.id
}

func (r *Recipe) Title() string {
	return r.title
}

func (r *Recipe) Description() string {
	return r.description
}

func (r *Recipe) ExternalLink() string {
	return r.externalLink
}
