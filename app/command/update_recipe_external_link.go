package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeExternalLink struct {
	RecipeID     string
	ExternalLink string
}

func (u UpdateRecipeExternalLink) GetCommandIDPayload() string {
	return u.RecipeID
}

type UpdateRecipeExternalLinkHandler struct {
	update domain.UpdateRecipeExternalLink
}

func NewUpdateRecipeExternalLink(recipeID string, link string) UpdateRecipeExternalLink {
	return UpdateRecipeExternalLink{
		RecipeID:     recipeID,
		ExternalLink: link,
	}
}
func NewUpdateRecipeExternalLinkHandler(update domain.UpdateRecipeExternalLink) UpdateRecipeExternalLinkHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeExternalLinkHandler")
	}

	return UpdateRecipeExternalLinkHandler{update: update}
}

func (h UpdateRecipeExternalLinkHandler) Handle(ctx context.Context, cmd UpdateRecipeExternalLink) error {
	rid, err := domain.NewRIDFromCmd(cmd)
	if err != nil {
		return fmt.Errorf("invalid RecipeID uuid: %s", cmd.RecipeID)
	}

	if !domain.CanUpdateExternalLink(cmd.ExternalLink) {
		return errors.New("external link is too long")
	}
	return h.update.UpdateRecipeExternalLink(ctx, rid, cmd.ExternalLink)
}
