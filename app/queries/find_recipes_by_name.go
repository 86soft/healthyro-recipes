package queries

import (
	"context"
	"errors"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type FindRecipesByName struct {
	Name string
}

type FindRecipesByNameHandler func(ctx context.Context, query FindRecipesByName) ([]core.Recipe, error)

func NewFindRecipesByNameHandler(find core.FindRecipesByName, logger zerolog.Logger) (FindRecipesByNameHandler, error) {
	if find == nil {
		return nil, errors.New("NewFindRecipesByName dependency is nil")
	}
	return func(ctx context.Context, query FindRecipesByName) ([]core.Recipe, error) {
		return find(ctx, query.Name)
	}, nil
}
