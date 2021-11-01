package command

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeExternalLink struct {
	RecipeID     domain.RID
	ExternalLink string
}

type UpdateRecipeExternalLinkHandler struct {
	update domain.UpdateRecipeExternalLink
	get    domain.GetRecipe
}

func NewUpdateRecipeExternalLink(recipeID domain.RID, link string) UpdateRecipeExternalLink {
	return UpdateRecipeExternalLink{
		RecipeID:     recipeID,
		ExternalLink: link,
	}
}
func NewUpdateRecipeExternalLinkHandler(update domain.UpdateRecipeExternalLink, get domain.GetRecipe) UpdateRecipeExternalLinkHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeExternalLinkHandler")
	}
	if get == nil {
		panic("nil get inside NewUpdateRecipeExternalLinkHandler")
	}

	return UpdateRecipeExternalLinkHandler{update: update, get: get}
}

func (h UpdateRecipeExternalLinkHandler) Handle(ctx context.Context, cmd UpdateRecipeExternalLink) error {
	if !domain.CanUpdateExternalLink(cmd.ExternalLink) {
		return errors.New("title is too long")
	}
	return h.update.UpdateRecipeExternalLink(ctx, cmd.RecipeID, cmd.ExternalLink)
}
