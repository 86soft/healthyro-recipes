package commands

import (
	"context"
	"errors"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type UpdateRecipeTitle struct {
	RecipeID c.ID[c.Recipe]
	Title    string
}

type UpdateRecipeTitleHandler func(ctx context.Context, cmd UpdateRecipeTitle) error

func NewUpdateRecipeTitleHandler(updateFn c.UpdateRecipeTitle, logger zerolog.Logger) (UpdateRecipeTitleHandler, error) {
	if updateFn == nil {
		return nil, errors.New("NewUpdateRecipeTitleHandler - updateFn dependency is nil")
	}

	return func(ctx context.Context, cmd UpdateRecipeTitle) error {
		if len(cmd.Title) > c.TitleLengthLimit {
			return c.ErrLengthLimitExceeded
		}
		return updateFn(ctx, cmd.RecipeID, cmd.Title)
	}, nil
}
