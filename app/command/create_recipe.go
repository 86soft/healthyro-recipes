package command

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
)

type CreateRecipe struct {
	Title        string
	Description  string
	ExternalLink string
}

func NewCreateRecipe(title, description, externalLink string) CreateRecipe {
	return CreateRecipe{
		Title:        title,
		Description:  description,
		ExternalLink: externalLink,
	}
}

type CreateRecipeHandler struct {
	repo domain.AddRecipe
}

func NewCreateRecipeHandler(addRecipe domain.AddRecipe) CreateRecipeHandler {
	if addRecipe == nil {
		panic("nil addRecipe inside NewCreateRecipeHandler")
	}
	return CreateRecipeHandler{repo: addRecipe}
}

func (h CreateRecipeHandler) Handle(ctx context.Context, cmd CreateRecipe) (recipeID domain.RID, err error) {
	recipe, err := domain.NewRecipe(cmd.Title, cmd.Description, cmd.ExternalLink)
	if err != nil {
		return domain.NilRID, err
	}
	return h.repo.AddRecipe(ctx, &recipe)
}
