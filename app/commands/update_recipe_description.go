package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeDescription struct {
	RecipeID    string
	Description string
}

type UpdateRecipeDescriptionHandler struct {
	updateDescriptionFn func(ctx context.Context, recipeID domain.RecipeID, description string) error
}

func NewUpdateRecipeDescription(recipeID string, description string) UpdateRecipeDescription {
	return UpdateRecipeDescription{
		RecipeID:    recipeID,
		Description: description,
	}
}
func NewUpdateRecipeDescriptionHandler(repo domain.UpdateRecipeDescription) UpdateRecipeDescriptionHandler {
	if repo == nil {
		panic("nil updateDescriptionFn inside NewUpdateRecipeDescriptionHandler")
	}
	return UpdateRecipeDescriptionHandler{updateDescriptionFn: repo.UpdateRecipeDescription}
}

func (h *UpdateRecipeDescriptionHandler) Handle(ctx context.Context, cmd UpdateRecipeDescription) error {
	if len(cmd.Description) > domain.DescriptionLengthLimit {
		return domain.ErrLengthLimitExceeded
	}
	id := domain.NewRecipeID(cmd.RecipeID)
	return h.updateDescriptionFn(ctx, id, cmd.Description)
}
