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
	t, _ := queries.NewGetRecipeByIdHandler(repo)
	return Queries{
		GetRecipeById: t,
		ListRecipes:   queries.NewListRecipesHandler(repo),
	}
}
