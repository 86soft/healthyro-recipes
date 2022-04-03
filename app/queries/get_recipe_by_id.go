package queries

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type GetRecipeById struct {
	RecipeID core.ID[core.Recipe]
}

type GetRecipeByIdHandler func(ctx context.Context, query GetRecipeById) (core.Recipe, error)

func NewGetRecipeByIdHandler(get core.GetRecipe, logger zerolog.Logger) (GetRecipeByIdHandler, error) {
	if get == nil {
		return nil, errors.New("NewGetRecipeByIdHandler - get dependency is nil")
	}

	return func(ctx context.Context, query GetRecipeById) (core.Recipe, error) {
		return get(ctx, query.RecipeID)
	}, nil
}
