package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type DeleteRecipe struct {
	RecipeID domain.RID
}

func NewDeleteRecipe(id domain.RID) DeleteRecipe {
	return DeleteRecipe{RecipeID: id}
}

type DeleteRecipeHandler struct {
	delete domain.DeleteRecipe
}

func NewDeleteRecipeHandler(delete domain.DeleteRecipe) DeleteRecipeHandler {
	if delete == nil {
		panic("nil delete inside NewDeleteRecipeHandler")
	}
	return DeleteRecipeHandler{delete: delete}
}
func (h DeleteRecipeHandler) Handle(ctx context.Context, cmd DeleteRecipe) error {
	return h.delete.DeleteRecipe(ctx, cmd.RecipeID)
}
