package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
)

type GetRecipeById struct {
	RecipeID string
}

func NewGetRecipeById(id string) GetRecipeById {
	return GetRecipeById{RecipeID: id}
}

type GetRecipeByIdHandler struct {
	getRecipeFn func(ctx context.Context, recipeID core.RecipeID) (core.Recipe, error)
}

func NewGetRecipeByIdHandler(repo core.Store) (GetRecipeByIdHandler, error) {
	if repo == nil {
		panic("nil getRecipeFn inside NewGetRecipeByIdHandler")
	}

	return GetRecipeByIdHandler{getRecipeFn: repo.GetRecipe}, nil
}

func (h GetRecipeByIdHandler) Handle(ctx context.Context, query GetRecipeById) (core.Recipe, error) {
	id := core.NewRecipeID(query.RecipeID)
	return h.getRecipeFn(ctx, id)
}
