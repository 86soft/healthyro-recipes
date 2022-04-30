package queries

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type FindRecipesByNameAndTags struct {
	Name string
	Tags []core.Tag
}
type FindRecipesByNameAndTagsHandler func(context.Context, FindRecipesByNameAndTags) ([]core.Recipe, error)

func NewFindRecipesByNameAndTagsHandler(
	findFn core.FindRecipesByNameAndTags,
	logger zerolog.Logger,
) (FindRecipesByNameAndTagsHandler, error) {
	if findFn == nil {
		return nil, errors.New("NewFindRecipesByNameAndTagsHandler - findFn implementation is nil")
	}

	return func(ctx context.Context, cmd FindRecipesByNameAndTags) ([]core.Recipe, error) {
		return findFn(ctx, cmd.Name, cmd.Tags)
	}, nil
}
