package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type ListRecipes struct {
}

func NewListRecipes() ListRecipes {
	return ListRecipes{}
}

type ListRecipesHandler struct {
	getRecipesFn func(ctx context.Context) ([]domain.Recipe, error)
}

func NewListRecipesHandler(repo domain.Repository) ListRecipesHandler {
	if repo == nil {
		panic("nil getRecipeFn inside NewListRecipesHandler")
	}
	return ListRecipesHandler{getRecipesFn: repo.GetRecipes}
}

func (h ListRecipesHandler) Handle(ctx context.Context) ([]domain.Recipe, error) {
	return h.getRecipesFn(ctx)
}
