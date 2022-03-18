package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/app"
	d "github.com/86soft/healthyro-recipes/domain"
	uuid "github.com/google/uuid"
	"github.com/rs/zerolog"
)

type AddRecipe struct {
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

type AddRecipeHandler struct {
	addRecipeFn func(ctx context.Context, newRecipe *d.Recipe) error
	log         zerolog.Logger
}

func NewAddRecipeHandler(db d.Store, logger zerolog.Logger) (AddRecipeHandler, error) {
	if db == nil {
		return AddRecipeHandler{}, &app.NilDependencyError{
			Name: "NewAddRecipeHandler - db",
		}
	}
	return AddRecipeHandler{
		addRecipeFn: db.AddRecipe,
		log:         logger,
	}, nil
}

func (h *AddRecipeHandler) Handle(ctx context.Context, cmd AddRecipe) (d.ID[d.Recipe], error) {
	resources := make([]d.Resource, 0, len(cmd.Resources))
	cmd.mapResources(resources)

	tags := make([]d.Tag, 0, len(cmd.Tags))
	cmd.mapTags(tags)

	recipe := d.Recipe{
		ID:          d.ID[d.Recipe]{ID: uuid.New().String()},
		Title:       cmd.Title,
		Description: cmd.Description,
		Resources:   resources,
		Tags:        tags,
	}
	err := h.addRecipeFn(ctx, &recipe)
	if err != nil {
		return d.ID[d.Recipe]{}, err
	}
	return recipe.ID, nil
}

func (c *AddRecipe) mapResources(res []d.Resource) {
	for _, r := range res {
		res = append(res, d.Resource{
			ID:    d.CreateID[d.Resource](),
			Name:  r.Name,
			Kind:  r.Kind,
			Value: r.Value,
		})
	}
}

func (c *AddRecipe) mapTags(tags []d.Tag) {
	for _, t := range c.Tags {
		tags = append(tags, d.Tag{
			ID:   d.CreateID[d.Tag](),
			Name: t,
		})
	}
}
