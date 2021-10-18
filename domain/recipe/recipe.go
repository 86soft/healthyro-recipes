package recipe

import (
	"errors"
)

type Recipe struct {
	recipeUUID   string
	title        string
	description  string
	externalLink string
}

func NewRecipe(title string, description string, externalLink string) (*Recipe, error) {
	if title == "" {
		return nil, errors.New("empty title is not allowed")
	}
	return &Recipe{
		title:        title,
		description:  description,
		externalLink: externalLink,
	}, nil
}

// UnmarshalRecipe is used only for unmarshalling Recipe from db
func UnmarshalRecipe(recipeUUID string, title string, description string, externalLink string) (*Recipe, error) {
	if title == "" {
		return nil, errors.New("empty title is not allowed")
	}
	return &Recipe{
		recipeUUID:   recipeUUID,
		title:        title,
		description:  description,
		externalLink: externalLink,
	}, nil
}

func (r *Recipe) UpdateTitle(title string) error {
	if len(title) > 100 {
		return errors.New("title to long, maximal limit: 100 chars")
	}
	r.title = title
	return nil
}

func (r *Recipe) UpdateDescription(description string) error {
	if len(description) > 5000 {
		return errors.New("description to long, maximal limit: 5000 chars")
	}
	r.description = description
	return nil
}

func (r *Recipe) UpdateExternalLink(externalLink string) error {
	if len(externalLink) > 2000 {
		return errors.New("externalLink to long, maximal limit: 2000 chars")
	}
	r.externalLink = externalLink
	return nil
}

func (r Recipe) RecipeUUID() string {
	return r.recipeUUID
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
