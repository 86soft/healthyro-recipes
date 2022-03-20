package commands

import (
	"context"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type CreateTag struct {
	Name string
}

type CreateTagHandler struct {
	createTagFn func(ctx context.Context, name string) (c.ID[c.Tag], error)
	log         zerolog.Logger
}

func NewCreateTagHandler(
	removeTag func(ctx context.Context, name string) (c.ID[c.Tag], error),
	logger zerolog.Logger,
) (CreateTagHandler, error) {
	if removeTag == nil {
		return CreateTagHandler{}, &c.NilDependencyError{
			Name: "RemoveTagFromRecipeHandler - createTagFn",
		}
	}
	return CreateTagHandler{
		createTagFn: removeTag,
		log:         logger,
	}, nil
}

func (h *CreateTagHandler) Handle(ctx context.Context, cmd CreateTag) (c.ID[c.Tag], error) {
	return h.createTagFn(ctx, cmd.Name)
}
