package commands

import (
	"context"
	"errors"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type UpdateRecipeDescription struct {
	RecipeID    d.ID[d.Recipe]
	Description string
}

type UpdateRecipeDescriptionHandler func(ctx context.Context, cmd UpdateRecipeDescription) error

func NewUpdateRecipeDescriptionHandler(
	updateFn d.UpdateRecipeDescription,
	logger l.Logger) (UpdateRecipeDescriptionHandler, error) {

	if updateFn == nil {
		return nil, errors.New("NewUpdateRecipeDescriptionHandler - updateFn dependency is nil")
	}
	return func(ctx context.Context, cmd UpdateRecipeDescription) error {
		if len(cmd.Description) > d.DescriptionLengthLimit {
			return d.ErrLengthLimitExceeded
		}
		return updateFn(ctx, cmd.RecipeID, cmd.Description)
	}, nil
}
