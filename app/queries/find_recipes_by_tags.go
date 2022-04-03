package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type FindRecipesByTags struct {
	Tags []core.Tag
}
type FindRecipesByTagsHandler func(context.Context, FindRecipesByTags) ([]core.Recipe, error)

func NewFindRecipesByTagsHandler(findFn core.FindRecipesByTags, logger zerolog.Logger) (FindRecipesByTagsHandler, error) {
	if findFn == nil {
		return nil, errors.New("NewFindRecipesByTags implementation is nil")
	}

	return func(ctx context.Context, cmd FindRecipesByTags) ([]core.Recipe, error) {
		return findFn(ctx, cmd.Tags)
	}, nil
}
