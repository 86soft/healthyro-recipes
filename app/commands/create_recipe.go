package commands

import (
	"context"
	"errors"
	d "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type CreateRecipe struct {
	Title       string
	Description string
	Resources   []RecipeResources
	Tags        []string
}

type RecipeResources struct {
	Name  string
	Kind  string
	Value string
}

type CreateRecipeHandler func(ctx context.Context, cmd CreateRecipe) (d.ID[d.Recipe], error)

func NewCreateRecipeHandler(
	createFn d.CreateRecipe,
	logger zerolog.Logger,
) (CreateRecipeHandler, error) {
	if createFn == nil {
		return nil, errors.New("NewCreateRecipeHandler - createFn dependency is nil")
	}
	return func(ctx context.Context, cmd CreateRecipe) (d.ID[d.Recipe], error) {
		resources := make([]d.Resource, len(cmd.Resources))
		cmd.mapResources(resources)

		tags := make([]d.Tag, len(cmd.Tags))
		cmd.mapTags(tags)

		recipe := d.Recipe{
			ID:          d.CreateID[d.Recipe](),
			Title:       cmd.Title,
			Description: cmd.Description,
			Resources:   resources,
			Tags:        tags,
		}
		err := createFn(ctx, &recipe)
		if err != nil {
			return d.ID[d.Recipe]{}, err
		}
		return recipe.ID, nil
	}, nil
}

func (c *CreateRecipe) mapResources(res []d.Resource) {
	for i, r := range c.Resources {
		res[i] = d.Resource{
			ID:    d.CreateID[d.Resource](),
			Name:  r.Name,
			Kind:  r.Kind,
			Value: r.Value,
		}
	}
}

func (c *CreateRecipe) mapTags(tags []d.Tag) {
	for i, t := range c.Tags {
		tags[i] = d.Tag{
			Name: t,
		}
	}
}
