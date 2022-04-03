package commands

import (
	"context"
	"errors"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type DeleteRecipe struct {
	RecipeID d.ID[d.Recipe]
}

type DeleteRecipeHandler func(ctx context.Context, cmd DeleteRecipe) error

func NewDeleteRecipeHandler(deleteFn d.DeleteRecipe, logger l.Logger) (DeleteRecipeHandler, error) {
	if deleteFn == nil {
		return nil, errors.New("NewDeleteRecipeHandler - deleteFn dependency is nil")
	}
	return func(ctx context.Context, cmd DeleteRecipe) error {
		return deleteFn(ctx, cmd.RecipeID)
	}, nil
}
