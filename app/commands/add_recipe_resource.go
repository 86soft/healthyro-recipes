package commands

import (
	"context"
	"errors"
	d "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type AddRecipeResource struct {
	Name     string
	Kind     string
	Value    string
	RecipeID d.ID[d.Recipe]
}

type AddRecipeResourceHandler func(ctx context.Context, cmd AddRecipeResource) error

func NewAddRecipeResourceHandler(add d.AddRecipeResource, logger zerolog.Logger) (AddRecipeResourceHandler, error) {
	if add == nil {
		return nil, errors.New("NewAddRecipeResourceHandler - add dependency is nil")
	}
	return func(ctx context.Context, cmd AddRecipeResource) error {
		r := d.Resource{
			ID:    d.CreateID[d.Resource](),
			Name:  cmd.Name,
			Kind:  cmd.Kind,
			Value: cmd.Value,
		}
		return add(ctx, cmd.RecipeID, &r)
	}, nil
}
