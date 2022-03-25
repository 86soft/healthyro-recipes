package commands

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type AddRecipeResource struct {
	Name     string
	Kind     string
	Value    string
	RecipeID d.ID[d.Recipe]
}

type AddRecipeResourceHandler struct {
	addRecipeResource func(
		ctx context.Context,
		id d.ID[d.Recipe],
		r *d.Resource,
	) error
	logger zerolog.Logger
}

func NewAddRecipeResourceHandler(fn func(
	ctx context.Context,
	id d.ID[d.Recipe],
	r *d.Resource,
) error, logger zerolog.Logger) (AddRecipeResourceHandler, error) {
	if fn == nil {
		return AddRecipeResourceHandler{}, &d.NilDependencyError{Name: "AddRecipeResourceHandler - fn"}
	}
	return AddRecipeResourceHandler{
		addRecipeResource: fn,
		logger:            logger,
	}, nil
}

func (h *AddRecipeResourceHandler) Handle(ctx context.Context, cmd AddRecipeResource) error {
	r := d.Resource{
		ID:    d.CreateID[d.Resource](),
		Name:  cmd.Name,
		Kind:  cmd.Kind,
		Value: cmd.Value,
	}
	return h.addRecipeResource(ctx, cmd.RecipeID, &r)
}
