package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type AddRecipeResource struct {
	name  string
	kind  string
	value string
	id    domain.RecipeID
}

type AddRecipeResourceHandler struct {
	addRecipeResource func(ctx context.Context, id domain.RecipeID, r *domain.Resource) error
}

func NewAddRecipeResource(name, kind, value string) AddRecipeResource {
	return AddRecipeResource{
		name:  name,
		kind:  kind,
		value: value,
	}
}

func NewAddRecipeResourceHandler(fn func(
	ctx context.Context,
	id domain.RecipeID,
	r *domain.Resource,
) error) (AddRecipeResourceHandler, error) {
	if fn == nil {
		return AddRecipeResourceHandler{}, &NilDependencyError{name: "AddRecipeResourceHandler"}
	}
	return AddRecipeResourceHandler{addRecipeResource: fn}, nil
}

func (h *AddRecipeResourceHandler) Handle(ctx context.Context, cmd AddRecipeResource) error {
	r := domain.Resource{
		Name:  cmd.name,
		Kind:  cmd.kind,
		Value: cmd.value,
	}
	return h.addRecipeResource(ctx, cmd.id, &r)
}
