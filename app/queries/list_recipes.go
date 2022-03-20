package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type ListRecipes struct {
}

func NewListRecipes() ListRecipes {
	return ListRecipes{}
}

type ListRecipesHandler struct {
	getRecipesFn func(ctx context.Context) ([]core.Recipe, error)
	logger       zerolog.Logger
}

func NewListRecipesHandler(repo func(ctx context.Context) ([]core.Recipe, error), logger zerolog.Logger) (ListRecipesHandler, error) {
	if repo == nil {
		return ListRecipesHandler{}, &core.NilDependencyError{Name: "NewListRecipesHandler - repo"}
	}
	return ListRecipesHandler{getRecipesFn: repo, logger: logger}, nil
}

func (h ListRecipesHandler) Handle(ctx context.Context) ([]core.Recipe, error) {
	return h.getRecipesFn(ctx)
}
