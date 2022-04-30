package commands

import (
	"context"
	"errors"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type AddTagToRecipe struct {
	Name         string
	TagID        c.ID[c.Tag]
	RecipeID     c.ID[c.Recipe]
	CreateNewTag bool
}

type AddTagToRecipeHandler func(ctx context.Context, cmd AddTagToRecipe) error

func NewAddTagToRecipeHandler(
	addTagFn c.AddTagToRecipe,
	createTagFn c.CreateTag,
	logger zerolog.Logger,
) (AddTagToRecipeHandler, error) {
	if addTagFn == nil {
		return nil, errors.New("NewAddTagToRecipeHandler - addTagFn dependency is nil")
	}
	if createTagFn == nil {
		return nil, errors.New("NewAddTagToRecipeHandler - createTagFn dependency is nil")
	}
	return func(ctx context.Context, cmd AddTagToRecipe) error {
		tag, err := getOrCreateTag(ctx, cmd, createTagFn)
		if err != nil {
			return err
		}
		return addTagFn(ctx, cmd.RecipeID, &tag)
	}, nil
}

func getOrCreateTag(ctx context.Context, cmd AddTagToRecipe, createTagFn c.CreateTag) (c.Tag, error) {
	var err error
	if cmd.CreateNewTag {
		_, err = createTagFn(ctx, cmd.Name)
	}
	return c.Tag{
		Name: cmd.Name,
	}, err
}
