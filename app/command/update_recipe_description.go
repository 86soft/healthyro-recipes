package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type UpdateRecipeDescription struct {
	RecipeUUID  string
	Description string
}

type UpdateRecipeDescriptionHandler struct {
	update recipe.UpdateRecipe
	get    recipe.GetRecipe
}

func NewUpdateRecipeDescriptionHandler(update recipe.UpdateRecipe, get recipe.GetRecipe) UpdateRecipeDescriptionHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeDescriptionHandler")
	}
	if get == nil {
		panic("nil get inside NewUpdateRecipeDescriptionHandler")
	}

	return UpdateRecipeDescriptionHandler{update: update, get: get}
}

func (h UpdateRecipeDescriptionHandler) Handle(ctx context.Context, cmd UpdateRecipeDescription) error {
	r, err := h.get.GetRecipe(ctx, cmd.RecipeUUID)
	if err != nil {
		return err
	}
	err = r.UpdateDescription(cmd.Description)
	if err != nil {
		return err
	}
	return h.update.UpdateRecipe(ctx, r)
}
