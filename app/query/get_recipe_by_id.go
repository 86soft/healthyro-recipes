package query

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type GetRecipeById struct {
	RecipeID string
}

func (g GetRecipeById) GetQueryIDPayload() string {
	return g.RecipeID
}

func NewGetRecipeById(id string) GetRecipeById {
	return GetRecipeById{RecipeID: id}
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
	rid, err := domain.NewRIDFromQuery(query)
	if err != nil {
		return domain.NilRecipe, err
	}
	return h.get.GetRecipe(ctx, rid)
}
