package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type GetRecipeById struct {
	RecipeID string
}

func NewGetRecipeById(id string) GetRecipeById {
	return GetRecipeById{RecipeID: id}
}

type GetRecipeByIdHandler struct {
	getRecipeFn func(ctx context.Context, recipeID domain.RecipeID) (domain.Recipe, error)
}

func NewGetRecipeByIdHandler(repo domain.GetRecipe) GetRecipeByIdHandler {
	if repo == nil {
		panic("nil getRecipeFn inside NewGetRecipeByIdHandler")
	}

	return GetRecipeByIdHandler{getRecipeFn: repo.GetRecipe}
}

func (h GetRecipeByIdHandler) Handle(ctx context.Context, query GetRecipeById) (domain.Recipe, error) {
	id := domain.NewRecipeID(query.RecipeID)
	return h.getRecipeFn(ctx, id)
}
