package commands

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type DeleteRecipeResource struct {
	RecipeID   d.ID[d.Recipe]
	ResourceID d.ID[d.Resource]
}

type RemoveRecipeResourceHandler struct {
	removeRecipeResource func(
		ctx context.Context,
		recipeID d.ID[d.Recipe],
		resourceID d.ID[d.Resource]) error
	log l.Logger
}

func NewRemoveRecipeResourceHandler(
	fn func(
		ctx context.Context,
		recipeID d.ID[d.Recipe],
		resourceID d.ID[d.Resource]) error,
	logger l.Logger,
) (RemoveRecipeResourceHandler, error) {
	if fn == nil {
		return RemoveRecipeResourceHandler{}, &d.NilDependencyError{Name: "RemoveRecipeResourceHandler"}
	}
	return RemoveRecipeResourceHandler{
		removeRecipeResource: fn,
		log:                  logger,
	}, nil
}

func (h *RemoveRecipeResourceHandler) Handle(ctx context.Context, cmd DeleteRecipeResource) error {
	return h.removeRecipeResource(ctx, cmd.RecipeID, cmd.ResourceID)
}
