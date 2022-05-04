package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type CreateRecipe struct {
	Title       string
	Description string
	Resources   []RecipeResources
	Tags        []CreateRecipeCmdTag
}

type RecipeResources struct {
	Name  string
	Kind  string
	Value string
}

type CreateRecipeCmdTag struct {
	Name   string
	Create bool
}

type CreateRecipeHandler func(ctx context.Context, cmd CreateRecipe) (core.ID[core.Recipe], error)

func NewCreateRecipeHandler(
	createRecipeFn core.CreateRecipe,
	createTagsFn core.CreateTags,
	addRecipeToTagsFn core.AddRecipeToTags,
	logger zerolog.Logger,
) (CreateRecipeHandler, error) {
	if createRecipeFn == nil {
		return nil, errors.New("NewCreateRecipeHandler - createRecipeFn dependency is nil")
	}
	if createTagsFn == nil {
		return nil, errors.New("NewCreateRecipeHandler - createTagsFn dependency is nil")
	}
	if addRecipeToTagsFn == nil {
		return nil, errors.New("NewCreateRecipeHandler - addRecipeToTagsFn dependency is nil")
	}
	return func(ctx context.Context, cmd CreateRecipe) (core.ID[core.Recipe], error) {
		recipeID := core.CreateID[core.Recipe]()

		resources := make([]core.Resource, len(cmd.Resources))
		cmd.mapResources(resources)
		allTags, newTags := cmd.mapTags()

		if len(newTags) > 0 {
			if err := createTagsFn(ctx, newTags); err != nil {
				return core.ID[core.Recipe]{}, fmt.Errorf("createTagsFn: %w", err)
			}
		}

		recipe := core.Recipe{
			ID:          recipeID,
			Title:       cmd.Title,
			Description: cmd.Description,
			Resources:   resources,
			Tags:        allTags,
		}

		err := createRecipeFn(ctx, &recipe)
		if err != nil {
			return core.ID[core.Recipe]{}, fmt.Errorf("createRecipeFn: %w", err)
		}

		err = addRecipeToTagsFn(ctx, recipeID, allTags)
		if err != nil {
			return core.ID[core.Recipe]{}, fmt.Errorf("addRecipeToTagsFn: %w", err)
		}
		return recipe.ID, nil
	}, nil
}

func (c *CreateRecipe) mapResources(res []core.Resource) {
	for i, r := range c.Resources {
		res[i] = core.Resource{
			ID:    core.CreateID[core.Resource](),
			Name:  r.Name,
			Kind:  r.Kind,
			Value: r.Value,
		}
	}
}

func (c *CreateRecipe) mapTags() (allTags []core.Tag, newTags []core.Tag) {
	allTags = make([]core.Tag, 0, len(c.Tags))
	for _, t := range c.Tags {
		tag := core.Tag{
			Name: t.Name,
		}
		if t.Create {
			newTags = append(newTags, tag)
		}
		allTags = append(allTags, tag)
	}
	return
}
