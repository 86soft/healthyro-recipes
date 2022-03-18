package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/domain"
	"github.com/google/uuid"
)

type DeleteRecipe struct {
	recipeID string
}

type DeleteRecipeHandler struct {
	deleteRecipeFn func(ctx context.Context, recipeID domain.ID[domain.Recipe]) error
}

func NewDeleteRecipe(id string) DeleteRecipe {
	return DeleteRecipe{recipeID: id}
}

func NewDeleteRecipeHandler(repo domain.Store) DeleteRecipeHandler {
	if repo == nil {
		panic("nil deleteRecipeFn inside NewDeleteRecipeHandler")
	}
	return DeleteRecipeHandler{deleteRecipeFn: repo.DeleteRecipe}
}
func (h *DeleteRecipeHandler) Handle(ctx context.Context, cmd DeleteRecipe) error {
	_, err := uuid.Parse(cmd.recipeID)
	if err != nil {
		return &app.CorruptedUUIDError{
			ID:      cmd.recipeID,
			Details: err.Error(),
		}
	}
	return h.deleteRecipeFn(ctx, id)
}
