package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type UpdateRecipeTitle struct {
	RecipeUUID string
	Title      string
}

type UpdateRecipeTitleHandler struct {
	update recipe.UpdateRecipe
	get    recipe.GetRecipe
}

func NewUpdateRecipeTitleHandler(update recipe.UpdateRecipe, get recipe.GetRecipe) UpdateRecipeTitleHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeTitleHandler")
	}
	if get == nil {
		panic("nil get inside NewUpdateRecipeTitleHandler")
	}

	return UpdateRecipeTitleHandler{update: update, get: get}
}

func (h UpdateRecipeTitleHandler) Handle(ctx context.Context, cmd UpdateRecipeTitle) error {
	rcp, err := h.get.GetRecipe(ctx, cmd.RecipeUUID)
	if err != nil {
		return err
	}
	err = rcp.UpdateTitle(cmd.Title)
	if err != nil {
		return err
	}
	return h.update.UpdateRecipe(ctx, rcp)
}
