package queries

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type ListRecipes struct {
}

type ListRecipesHandler func(ctx context.Context) ([]core.Recipe, error)

func NewListRecipesHandler(list core.ListRecipes, logger zerolog.Logger) (ListRecipesHandler, error) {
	if list == nil {
		return nil, errors.New("NewListRecipesHandler - list dependency is nil")
	}
	return func(ctx context.Context) ([]core.Recipe, error) {
		return list(ctx)
	}, nil
}
