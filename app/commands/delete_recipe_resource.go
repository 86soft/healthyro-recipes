package commands

import (
	"context"
	"errors"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type RemoveResourceFromRecipe struct {
	RecipeID   d.ID[d.Recipe]
	ResourceID d.ID[d.Resource]
}

type RemoveResourceFromRecipeHandler func(ctx context.Context, cmd RemoveResourceFromRecipe) error

func NewRemoveResourceFromRecipeHandler(
	removeResourceFn d.RemoveResourceFromRecipe,
	logger l.Logger,
) (RemoveResourceFromRecipeHandler, error) {
	if removeResourceFn == nil {
		return nil, errors.New("NewRemoveResourceFromRecipeHandler - removeResourceFn dependency is nil")
	}
	return func(ctx context.Context, cmd RemoveResourceFromRecipe) error {
		return removeResourceFn(ctx, cmd.RecipeID, cmd.ResourceID)
	}, nil
}
