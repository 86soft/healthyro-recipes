package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type FindRecipesByTags struct {
	Tags []core.Tag
}

type FindRecipesByTagsHandler struct {
	findFn func(ctx context.Context, tags []core.Tag) ([]core.Recipe, error)
	logger zerolog.Logger
}

func NewFindRecipesByTagsHandler(
	fn func(ctx context.Context, tags []core.Tag) ([]core.Recipe, error),
	logger zerolog.Logger) (FindRecipesByTagsHandler, error) {
	if fn == nil {
		return FindRecipesByTagsHandler{}, &core.NilDependencyError{Name: "FindRecipesByTagsHandler - fn"}
	}

	return FindRecipesByTagsHandler{findFn: fn, logger: logger}, nil
}

func (h *FindRecipesByTagsHandler) Handle(ctx context.Context, query FindRecipesByTags) ([]core.Recipe, error) {
	return h.findFn(ctx, query.Tags)
}
