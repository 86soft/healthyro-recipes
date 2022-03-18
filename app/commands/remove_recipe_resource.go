package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type RemoveRecipeResource struct {
	name  string
	kind  string
	value string
}

type RemoveRecipeResourceHandler struct {
	removeRecipeResource func(ctx context.Context, id domain.RecipeID, r *domain.Resource) error
}

func NewRemoveRecipeResource(name, kind, value string) RemoveRecipeResource {
	return RemoveRecipeResource{
		name:  name,
		kind:  kind,
		value: value,
	}
}

/*func NewRemoveRecipeResourceHandler(
	fn func(
		ctx context.Context,
		ID domain.RecipeID,
		Name string) error,
) (RemoveRecipeResourceHandler, error) {
	if fn == nil {
		return RemoveRecipeResourceHandler{}, &NilDependencyError{Name: "RemoveRecipeResourceHandler"}
	}
	return RemoveRecipeResourceHandler{removeRecipeResource: fn}, nil
}

func (h *RemoveRecipeResourceHandler) Handle(ctx context.Context, cmd AddRecipe) error {
	recipe, err := domain.NewRecipe(cmd.Title, cmd.Description)
	if err != nil {
		return err
	}
	return h.removeRecipeResource(ctx, &recipe)
}*/
