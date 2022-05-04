package commands

import (
	"context"
	"errors"
	"fmt"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type AddTagToRecipe struct {
	Name         string
	RecipeID     c.ID[c.Recipe]
	CreateNewTag bool
}

type AddTagToRecipeHandler func(ctx context.Context, cmd AddTagToRecipe) error

func NewAddTagToRecipeHandler(
	addTagToRecipeFn c.AddTagToRecipe,
	createTagFn c.CreateTag,
	checkRecipeForTagFn c.CheckRecipeForTag,
	tagExistFn c.TagExist,
	logger zerolog.Logger,
) (AddTagToRecipeHandler, error) {
	if addTagToRecipeFn == nil {
		return nil, errors.New("NewAddTagToRecipeHandler - addTagToRecipeFn dependency is nil")
	}
	if createTagFn == nil {
		return nil, errors.New("NewAddTagToRecipeHandler - createTagFn dependency is nil")
	}
	if checkRecipeForTagFn == nil {
		return nil, errors.New("NewAddTagToRecipeHandler - checkRecipeForTagFn dependency is nil")
	}
	if tagExistFn == nil {
		return nil, errors.New("NewAddTagToRecipeHandler - tagExistFn dependency is nil")
	}
	return func(ctx context.Context, cmd AddTagToRecipe) error {
		tag, err := getOrCreateTag(ctx, cmd, createTagFn)
		if err != nil {
			return err
		}

		exist, err := tagExistFn(ctx, tag.Name)
		if err != nil {
			return fmt.Errorf("tagExistFn: %w", err)
		}
		if !exist {
			return fmt.Errorf("tag: %s doesn't exist", tag.Name)
		}

		hasTag, err := checkRecipeForTagFn(ctx, cmd.RecipeID, tag.Name)
		if err != nil {
			return fmt.Errorf("checkRecipeForTagFn: %w", err)
		}
		if hasTag {
			return fmt.Errorf("recipeID: %s already have tag: %s", cmd.RecipeID.Value.String(), tag.Name)
		}

		return addTagToRecipeFn(ctx, cmd.RecipeID, &tag)
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
