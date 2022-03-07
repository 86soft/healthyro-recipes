package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type CreateRecipe struct {
	title       string
	description string
}

type CreateRecipeHandler struct {
	createRecipe func(ctx context.Context, newRecipe *domain.Recipe) error
}

func NewCreateRecipe(title, description string) CreateRecipe {
	return CreateRecipe{
		title:       title,
		description: description,
	}
}

func NewCreateRecipeHandler(db domain.Repository) CreateRecipeHandler {
	if db == nil {
		panic("nil db inside NewCreateRecipeHandler")
	}
	return CreateRecipeHandler{createRecipe: db.AddRecipe}
}

func (h *CreateRecipeHandler) Handle(ctx context.Context, cmd CreateRecipe) error {
	recipe, err := domain.NewRecipe(cmd.title, cmd.description)
	if err != nil {
		return err
	}
	return h.createRecipe(ctx, &recipe)
}
