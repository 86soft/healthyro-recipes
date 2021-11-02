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

func NewRecipe(title string, description string, externalLink string) (Recipe, error) {
	if title == "" {
		return NilRecipe, errors.New("empty title is not allowed")
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
	if CanUpdateTitle(title) {
		return errors.New("title to long, maximal limit: 100 chars")
	}
	r.title = title
	return nil
}

func (r *Recipe) UpdateDescription(description string) error {
	if CanUpdateDescription(description) {
		return errors.New("description to long, maximal limit: 5000 chars")
	}
	r.description = description
	return nil
}

func (r *Recipe) UpdateExternalLink(externalLink string) error {
	if !CanUpdateExternalLink(externalLink) {
		return errors.New("externalLink to long, maximal limit: 2000 chars")
	}
	r.externalLink = externalLink
	return nil
}

func (r Recipe) RecipeID() uuid.UUID {
	return r.id
}
func (r Recipe) Title() string {
	return r.title
}
func (r Recipe) Description() string {
	return r.description
}
func (r Recipe) ExternalLink() string {
	return r.externalLink
}
