package commands

import (
	"context"
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

type CreateRecipeHandler struct {
	createRecipeFn func(ctx context.Context, newRecipe *d.Recipe) error
	log            zerolog.Logger
}

func NewCreateRecipeHandler(
	fn func(ctx context.Context, r *d.Recipe) error,
	logger zerolog.Logger,
) (CreateRecipeHandler, error) {
	if fn == nil {
		return CreateRecipeHandler{}, &d.NilDependencyError{
			Name: "NewCreateRecipeHandler - fn",
		}
	}
	return CreateRecipeHandler{
		createRecipeFn: fn,
		log:            logger,
	}, nil
}

func (h *CreateRecipeHandler) Handle(ctx context.Context, cmd CreateRecipe) (d.ID[d.Recipe], error) {
	resources := make([]d.Resource, 0, len(cmd.Resources))
	cmd.mapResources(resources)

	tags := make([]d.Tag, 0, len(cmd.Tags))
	cmd.mapTags(tags)

	recipe := d.Recipe{
		ID:          d.CreateID[d.Recipe](),
		Title:       cmd.Title,
		Description: cmd.Description,
		Resources:   resources,
		Tags:        tags,
	}
	err := h.createRecipeFn(ctx, &recipe)
	if err != nil {
		return d.ID[d.Recipe]{}, err
	}
	return recipe.ID, nil
}

func (c *CreateRecipe) mapResources(res []d.Resource) {
	for _, r := range res {
		res = append(res, d.Resource{
			ID:    d.CreateID[d.Resource](),
			Name:  r.Name,
			Kind:  r.Kind,
			Value: r.Value,
		})
	}
}

func (c *CreateRecipe) mapTags(tags []d.Tag) {
	for _, t := range c.Tags {
		tags = append(tags, d.Tag{
			ID:   d.CreateID[d.Tag](),
			Name: t,
		})
	}
}
