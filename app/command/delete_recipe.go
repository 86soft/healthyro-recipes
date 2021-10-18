package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type DeleteRecipe struct {
	RecipeUUID string
}

type DeleteRecipeHandler struct {
	delete recipe.DeleteRecipe
}

func NewDeleteRecipeHandler(delete recipe.DeleteRecipe) DeleteRecipeHandler {
	if delete == nil {
		panic("nil delete inside NewDeleteRecipeHandler")
	}
	return DeleteRecipeHandler{delete: delete}
}
func (h DeleteRecipeHandler) Handle(ctx context.Context, cmd DeleteRecipe) error {
	return h.delete.DeleteRecipe(ctx, cmd.RecipeUUID)
}
