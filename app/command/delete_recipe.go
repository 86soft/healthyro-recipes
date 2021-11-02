package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type DeleteRecipe struct {
	RecipeID string
}

func (d DeleteRecipe) GetCommandIDPayload() string {
	return d.RecipeID
}

func NewDeleteRecipe(id string) DeleteRecipe {
	return DeleteRecipe{RecipeID: id}
}

type DeleteRecipeHandler struct {
	delete domain.DeleteRecipe
}

func NewDeleteRecipeHandler(delete domain.DeleteRecipe) DeleteRecipeHandler {
	if delete == nil {
		panic("nil delete inside NewDeleteRecipeHandler")
	}
	return DeleteRecipeHandler{delete: delete}
}
func (h DeleteRecipeHandler) Handle(ctx context.Context, cmd DeleteRecipe) error {
	rid, err := domain.NewRIDFromCmd(cmd)
	if err != nil {
		return err
	}
	return h.delete.DeleteRecipe(ctx, rid)
}
