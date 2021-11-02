package command

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeDescription struct {
	RecipeID    string
	Description string
}

func (u UpdateRecipeDescription) GetCommandIDPayload() string {
	return u.RecipeID
}

type UpdateRecipeDescriptionHandler struct {
	update domain.UpdateRecipeDescription
}

func NewUpdateRecipeDescription(recipeID string, description string) UpdateRecipeDescription {
	return UpdateRecipeDescription{
		RecipeID:    recipeID,
		Description: description,
	}
}
func NewUpdateRecipeDescriptionHandler(update domain.UpdateRecipeDescription) UpdateRecipeDescriptionHandler {
	if update == nil {
		panic("nil update inside NewUpdateRecipeDescriptionHandler")
	}
	return UpdateRecipeDescriptionHandler{update: update}
}

func (h UpdateRecipeDescriptionHandler) Handle(ctx context.Context, cmd UpdateRecipeDescription) error {
	rid, err := domain.NewRIDFromCmd(cmd)
	if err != nil {
		return err
	}

	if !domain.CanUpdateDescription(cmd.Description) {
		return errors.New("description is too long")
	}
	return h.update.UpdateRecipeDescription(ctx, rid, cmd.Description)
}
