package query

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type ListRecipes struct {
}
type ListRecipesHandler struct {
	get domain.GetRecipes
}

func NewListRecipesHandler(get domain.GetRecipes) ListRecipesHandler {
	if get == nil {
		panic("nil get inside NewListRecipesHandler")
	}
	return ListRecipesHandler{get: get}
}

func (h ListRecipesHandler) Handle(ctx context.Context) ([]domain.Recipe, error) {
	return h.get.GetRecipes(ctx)
}
