package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type DeleteRecipe struct {
	recipeID string
}

type DeleteRecipeHandler struct {
	deleteRecipeFn func(ctx context.Context, recipeID domain.RecipeID) error
}

func NewDeleteRecipe(id string) DeleteRecipe {
	return DeleteRecipe{recipeID: id}
}

func NewDeleteRecipeHandler(repo domain.Repository) DeleteRecipeHandler {
	if repo == nil {
		panic("nil deleteRecipeFn inside NewDeleteRecipeHandler")
	}
	return DeleteRecipeHandler{deleteRecipeFn: repo.DeleteRecipe}
}
func (h *DeleteRecipeHandler) Handle(ctx context.Context, cmd DeleteRecipe) error {
	id := domain.NewRecipeID(cmd.recipeID)
	return h.deleteRecipeFn(ctx, id)
}
