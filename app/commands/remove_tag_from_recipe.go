package commands

import (
	"context"
	"errors"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type RemoveTagFromRecipe struct {
	TagID    c.ID[c.Tag]
	RecipeID c.ID[c.Recipe]
}

type RemoveTagFromRecipeHandler func(ctx context.Context, cmd RemoveTagFromRecipe) error

func NewRemoveTagFromRecipeHandler(
	removeTagFn c.RemoveTagFromRecipe,
	logger zerolog.Logger,
) (RemoveTagFromRecipeHandler, error) {
	if removeTagFn == nil {
		return nil, errors.New("NewRemoveTagFromRecipeHandler - removeTagFn dependency is nil")
	}
	return func(ctx context.Context, cmd RemoveTagFromRecipe) error {
		return removeTagFn(ctx, cmd.RecipeID, cmd.TagID)
	}, nil
}
