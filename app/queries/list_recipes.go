package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
)

type ListRecipes struct {
}

func NewListRecipes() ListRecipes {
	return ListRecipes{}
}

type ListRecipesHandler struct {
	getRecipesFn func(ctx context.Context) ([]core.Recipe, error)
}

func NewListRecipesHandler(repo core.Store) ListRecipesHandler {
	if repo == nil {
		panic("nil getRecipeFn inside NewListRecipesHandler")
	}
	return ListRecipesHandler{getRecipesFn: repo.GetRecipes}
}

func (h ListRecipesHandler) Handle(ctx context.Context) ([]core.Recipe, error) {
	return h.getRecipesFn(ctx)
}
