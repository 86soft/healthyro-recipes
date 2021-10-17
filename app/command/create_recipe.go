package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type CreateRecipe struct {
	Title        string
	Description  string
	ExternalLink string
}

type CreateRecipeHandler struct {
	repo recipe.AddRecipe
}

func NewCreateRecipeHandler(addRecipe recipe.AddRecipe) CreateRecipeHandler {
	if addRecipe == nil {
		panic("nil addRecipe inside NewCreateRecipeHandler")
	}
	return CreateRecipeHandler{repo: addRecipe}
}

func (h CreateRecipeHandler) Handle(ctx context.Context, cmd CreateRecipe) (err error) {
	return h.repo.AddRecipe(ctx, cmd.Title, cmd.Description, cmd.ExternalLink)
}
