package commands

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type DeleteRecipe struct {
	RecipeID d.ID[d.Recipe]
}

type DeleteRecipeHandler struct {
	deleteRecipeFn func(ctx context.Context, id d.ID[d.Recipe]) error
	logger         l.Logger
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
	return h.deleteRecipeFn(ctx, cmd.RecipeID)
}
