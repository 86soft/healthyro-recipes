package commands

import (
	"context"
	"errors"
	"fmt"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type DeleteRecipe struct {
	RecipeID d.ID[d.Recipe]
}

type DeleteRecipeHandler func(ctx context.Context, cmd DeleteRecipe) error

func NewDeleteRecipeHandler(
	deleteFn d.DeleteRecipe,
	getRecipeFn d.GetRecipe,
	logger l.Logger) (DeleteRecipeHandler, error) {
	if deleteFn == nil {
		return nil, errors.New("NewDeleteRecipeHandler - deleteFn dependency is nil")
	}
	if getRecipeFn == nil {
		return nil, errors.New("NewDeleteRecipeHandler - getRecipeFn dependency is nil")
	}
	return func(ctx context.Context, cmd DeleteRecipe) error {
		r, err := getRecipeFn(ctx, cmd.RecipeID)
		if err != nil {
			return fmt.Errorf("getRecipeFn: %w", err)
		}
		return deleteFn(ctx, &r)
	}, nil
}
