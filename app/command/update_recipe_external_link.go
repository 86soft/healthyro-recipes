package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type UpdateRecipeExternalLink struct {
	RecipeUUID   string
	ExternalLink string
}

type UpdateRecipeExternalLinkHandler struct {
	update recipe.UpdateRecipe
	get    recipe.GetRecipe
}

func NewUpdateRecipeExternalLinkHandler(update recipe.UpdateRecipe, get recipe.GetRecipe) UpdateRecipeExternalLinkHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeExternalLinkHandler")
	}
	if get == nil {
		panic("nil get inside NewUpdateRecipeExternalLinkHandler")
	}

	return UpdateRecipeExternalLinkHandler{update: update, get: get}
}

func (h UpdateRecipeExternalLinkHandler) Handle(ctx context.Context, cmd UpdateRecipeExternalLink) error {
	rcp, err := h.get.GetRecipe(ctx, cmd.RecipeUUID)
	if err != nil {
		return err
	}
	err = rcp.UpdateExternalLink(cmd.ExternalLink)
	if err != nil {
		return err
	}
	return h.update.UpdateRecipe(ctx, rcp)
}
