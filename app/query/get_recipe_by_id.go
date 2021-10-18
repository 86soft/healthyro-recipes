package query

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type GetRecipeById struct {
	RecipeUUID string
}
type GetRecipeByIdHandler struct {
	get recipe.GetRecipe
}

func NewGetRecipeByIdHandler(get recipe.GetRecipe) GetRecipeByIdHandler {
	if get == nil {
		panic("nil get inside NewGetRecipeByIdHandler")
	}

	return GetRecipeByIdHandler{get: get}
}

func (h GetRecipeByIdHandler) Handle(ctx context.Context, query GetRecipeById) (recipe.Recipe, error) {
	return h.get.GetRecipe(ctx, query.RecipeUUID)
}
