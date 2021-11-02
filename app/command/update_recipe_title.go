package command

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeTitle struct {
	RecipeID string
	Title    string
}

func (u UpdateRecipeTitle) GetCommandIDPayload() string {
	return u.RecipeID
}

type UpdateRecipeTitleHandler struct {
	update domain.UpdateRecipeTitle
}

func NewUpdateRecipeTitle(recipeID string, title string) UpdateRecipeTitle {
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
	rid, err := domain.NewRIDFromCmd(cmd)
	if err != nil {
		return err
	}

	if !domain.CanUpdateTitle(cmd.Title) {
		return errors.New("title is too long")
	}
	return h.update.UpdateRecipeTitle(ctx, rid, cmd.Title)
}
