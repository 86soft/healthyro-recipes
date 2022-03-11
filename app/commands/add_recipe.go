package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
	uuid "github.com/google/uuid"
)

type AddRecipe struct {
	title       string
	description string
}

type AddRecipeHandler struct {
	addRecipeFn func(ctx context.Context, newRecipe *domain.Recipe) error
}

func NewAddRecipe(title, description string) AddRecipe {
	return AddRecipe{
		title:       title,
		description: description,
	}
}

func NewCreateRecipeHandler(db domain.Repository) AddRecipeHandler {
	if db == nil {
		panic("nil db inside NewCreateRecipeHandler")
	}
	return AddRecipeHandler{addRecipeFn: db.AddRecipe}
}

func (h *AddRecipeHandler) Handle(ctx context.Context, cmd AddRecipe) error {
	id := domain.NewRecipeID(uuid.New().String())
	recipe, err := domain.NewRecipe(id, cmd.title, cmd.description)
	if err != nil {
		return err
	}
	return h.addRecipeFn(ctx, &recipe)
}
