package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type FindRecipesByName struct {
	Name string
}

type FindRecipesByNameHandler struct {
	findFn func(ctx context.Context, name string) ([]core.Recipe, error)
	logger zerolog.Logger
}

func NewFindRecipesByNameHandler(
	fn func(ctx context.Context, name string) ([]core.Recipe, error),
	logger zerolog.Logger,
) (FindRecipesByNameHandler, error) {
	if fn == nil {
		return FindRecipesByNameHandler{}, &core.NilDependencyError{Name: "FindRecipesByNameHandler - fn"}
	}

	return FindRecipesByNameHandler{findFn: fn, logger: logger}, nil
}

func (h *FindRecipesByNameHandler) Handle(ctx context.Context, query FindRecipesByName) ([]core.Recipe, error) {
	return h.findFn(ctx, query.Name)
}
