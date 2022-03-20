package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type GetRecipeById struct {
	RecipeID core.ID[core.Recipe]
}

type GetRecipeByIdHandler struct {
	getRecipeFn func(ctx context.Context, id core.ID[core.Recipe]) (core.Recipe, error)
	logger      zerolog.Logger
}

func NewGetRecipeByIdHandler(fn func(ctx context.Context, id core.ID[core.Recipe]) (core.Recipe, error), logger zerolog.Logger) (GetRecipeByIdHandler, error) {
	if fn == nil {
		return GetRecipeByIdHandler{}, &core.NilDependencyError{Name: "NewGetRecipeByIdHandler - fn"}
	}

	return GetRecipeByIdHandler{getRecipeFn: fn, logger: logger}, nil
}

func (h GetRecipeByIdHandler) Handle(ctx context.Context, query GetRecipeById) (core.Recipe, error) {
	return h.getRecipeFn(ctx, query.RecipeID)
}
