package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeTitle struct {
	RecipeID string
	Title    string
}

type UpdateRecipeTitleHandler struct {
	updateRecipeFn func(ctx context.Context, recipeID domain.RecipeID, title string) error
}

func NewUpdateRecipeTitle(recipeID string, title string) UpdateRecipeTitle {
	return UpdateRecipeTitle{
		RecipeID: recipeID,
		Title:    title,
	}
}
func NewUpdateRecipeTitleHandler(repo domain.UpdateRecipeTitle) UpdateRecipeTitleHandler {
	if repo == nil {
		panic("nil updateDescriptionFn inside NewUpdateRecipeTitleHandler")
	}

	return UpdateRecipeTitleHandler{updateRecipeFn: repo.UpdateRecipeTitle}
}

func (h *UpdateRecipeTitleHandler) Handle(ctx context.Context, cmd UpdateRecipeTitle) error {
	id := domain.NewRecipeID(cmd.RecipeID)

	if len(cmd.Title) > domain.TitleLengthLimit {
		return domain.ErrLengthLimitExceeded
	}

	return h.updateRecipeFn(ctx, id, cmd.Title)
}
