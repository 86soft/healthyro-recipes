package app

import (
	"github.com/86soft/healthyro-recipes/app/queries"
	"github.com/86soft/healthyro-recipes/domain"
)

type Queries struct {
	GetRecipeById queries.GetRecipeByIdHandler
	ListRecipes   queries.ListRecipesHandler
}

func NewQueryHandlers(repo domain.Repository) Queries {
	return Queries{
		GetRecipeById: queries.NewGetRecipeByIdHandler(repo),
		ListRecipes:   queries.NewListRecipesHandler(repo),
	}
}
