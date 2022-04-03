package commands

import (
	"context"
	"errors"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type CreateTag struct {
	Name string
}

type CreateTagHandler func(ctx context.Context, cmd CreateTag) (c.ID[c.Tag], error)

func NewCreateTagHandler(createTagFn c.CreateTag, logger zerolog.Logger) (CreateTagHandler, error) {
	if createTagFn == nil {
		return nil, errors.New("NewCreateTagHandler - createTagFn dependency is nil")
	}
	return func(ctx context.Context, cmd CreateTag) (c.ID[c.Tag], error) {
		return createTagFn(ctx, cmd.Name)
	}, nil
}
