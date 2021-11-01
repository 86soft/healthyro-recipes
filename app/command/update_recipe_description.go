package command

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/domain"
)

type UpdateRecipeDescription struct {
	RecipeID    domain.RID
	Description string
}

type UpdateRecipeDescriptionHandler struct {
	update domain.UpdateRecipeDescription
}

func NewUpdateRecipeDescription(recipeID domain.RID, description string) UpdateRecipeDescription {
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
	if !domain.CanUpdateDescription(cmd.Description) {
		return errors.New("title is too long")
	}
	return h.update.UpdateRecipeDescription(ctx, cmd.RecipeID, cmd.Description)
}
