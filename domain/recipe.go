package domain

import (
	"errors"
)

type Recipe struct {
	id          RecipeID
	title       string
	description string
}

var (
	ErrLengthLimitExceeded = errors.New("length limit exceeded")
	ErrEmptyTitle          = errors.New("empty title is not allowed")
)

func NewRecipe(title string, description string) (Recipe, error) {
	if title == "" {
		return Recipe{}, ErrEmptyTitle
	}
	return Recipe{
		title:       title,
		description: description,
	}, nil
}

// UnmarshalRecipe is used only for unmarshalling Recipe from db
func UnmarshalRecipe(id RecipeID, title string, description string) Recipe {
	return Recipe{
		id:          id,
		title:       title,
		description: description,
	}
}

func (r *Recipe) UpdateTitle(title string) error {
	if len(title) > TitleLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.title = title
	return nil
}

func (r *Recipe) UpdateDescription(description string) error {
	if len(description) > DescriptionLengthLimit {
		return ErrLengthLimitExceeded
	}
	r.description = description
	return nil
}

func (r *Recipe) ID() string {
	return r.id.id
}

func (r *Recipe) Title() string {
	return r.title
}

func (r *Recipe) Description() string {
	return r.description
}
