package query

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain/recipe"
)

type ListRecipes struct {
}
type ListRecipesHandler struct {
	get recipe.GetRecipes
}

func NewListRecipesHandler(get recipe.GetRecipes) ListRecipesHandler {
	if get == nil {
		panic("nil get inside NewListRecipesHandler")
	}
	return ListRecipesHandler{get: get}
}

func (h ListRecipesHandler) Handle(ctx context.Context) ([]recipe.Recipe, error) {
	return h.get.GetRecipes(ctx)
}
