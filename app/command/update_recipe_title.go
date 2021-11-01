package command

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeTitle struct {
	RecipeID domain.RID
	Title    string
}

type UpdateRecipeTitleHandler struct {
	update domain.UpdateRecipeTitle
}

func NewUpdateRecipeTitle(recipeID domain.RID, title string) UpdateRecipeTitle {
	return UpdateRecipeTitle{
		RecipeID: recipeID,
		Title:    title,
	}
}
func NewUpdateRecipeTitleHandler(update domain.UpdateRecipeTitle) UpdateRecipeTitleHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeTitleHandler")
	}

	return UpdateRecipeTitleHandler{update: update}
}

func (h UpdateRecipeTitleHandler) Handle(ctx context.Context, cmd UpdateRecipeTitle) error {
	if !domain.CanUpdateTitle(cmd.Title) {
		return errors.New("title is too long")
	}
	return h.update.UpdateRecipeTitle(ctx, cmd.RecipeID, cmd.Title)
}
