package query

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type GetRecipeById struct {
	RecipeID domain.RID
}
type GetRecipeByIdHandler struct {
	get domain.GetRecipe
}

func NewGetRecipeByIdHandler(get domain.GetRecipe) GetRecipeByIdHandler {
	if get == nil {
		panic("nil get inside NewGetRecipeByIdHandler")
	}

	return GetRecipeByIdHandler{get: get}
}

func (h GetRecipeByIdHandler) Handle(ctx context.Context, query GetRecipeById) (domain.Recipe, error) {
	return h.get.GetRecipe(ctx, query.RecipeID)
}
