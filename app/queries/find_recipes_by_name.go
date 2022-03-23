package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type GetRecipesByName struct {
	RecipeID core.ID[core.Recipe]
}

type GetRecipesByNameHandler struct {
	getRecipeFn func(ctx context.Context, id core.ID[core.Recipe]) (core.Recipe, error)
	logger      zerolog.Logger
}

func NewGetRecipesByName(fn func(ctx context.Context, id core.ID[core.Recipe]) (core.Recipe, error), logger zerolog.Logger) (GetRecipeByIdHandler, error) {
	if fn == nil {
		return GetRecipeByIdHandler{}, &core.NilDependencyError{Name: "NewGetRecipeByIdHandler - fn"}
	}

	return GetRecipeByIdHandler{getRecipeFn: fn, logger: logger}, nil
}

func (h GetRecipesByNameHandler) Handle(ctx context.Context, query GetRecipesByName) (core.Recipe, error) {
	return h.getRecipeFn(ctx, query.RecipeID)
}
