package commands

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	"github.com/google/uuid"
	l "github.com/rs/zerolog"
)

type DeleteRecipe struct {
	recipeID string
}

type DeleteRecipeHandler struct {
	deleteRecipeFn func(ctx context.Context, id d.ID[d.Recipe]) error
	logger         l.Logger
}

func NewDeleteRecipe(id string) DeleteRecipe {
	return DeleteRecipe{recipeID: id}
}

func NewDeleteRecipeHandler(fn func(
	ctx context.Context,
	id d.ID[d.Recipe],
) error, logger l.Logger) (DeleteRecipeHandler, error) {
	if fn == nil {
		return DeleteRecipeHandler{}, &d.NilDependencyError{Name: "DeleteRecipeHandler - fn"}
	}
	return DeleteRecipeHandler{deleteRecipeFn: fn, logger: logger}, nil
}
func (h *DeleteRecipeHandler) Handle(ctx context.Context, cmd DeleteRecipe) error {
	_, err := uuid.Parse(cmd.recipeID)
	if err != nil {
		return &d.CorruptedUUIDError{
			ID:      cmd.recipeID,
			Details: err.Error(),
		}
	}
	return h.deleteRecipeFn(ctx, d.FromStringID[d.Recipe](cmd.recipeID))
}
