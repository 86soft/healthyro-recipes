package commands

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type RemoveResourceFromRecipe struct {
	RecipeID   d.ID[d.Recipe]
	ResourceID d.ID[d.Resource]
}

type RemoveResourceFromRecipeHandler struct {
	removeRecipeResource func(
		ctx context.Context,
		recipeID d.ID[d.Recipe],
		resourceID d.ID[d.Resource]) error
	log l.Logger
}

func NewRemoveResourceFromRecipeHandler(
	fn func(
		ctx context.Context,
		recipeID d.ID[d.Recipe],
		resourceID d.ID[d.Resource]) error,
	logger l.Logger,
) (RemoveResourceFromRecipeHandler, error) {
	if fn == nil {
		return RemoveResourceFromRecipeHandler{}, &d.NilDependencyError{Name: "RemoveResourceFromRecipe"}
	}
	return RemoveResourceFromRecipeHandler{
		removeRecipeResource: fn,
		log:                  logger,
	}, nil
}

func (h *RemoveResourceFromRecipeHandler) Handle(ctx context.Context, cmd RemoveResourceFromRecipe) error {
	return h.removeRecipeResource(ctx, cmd.RecipeID, cmd.ResourceID)
}
