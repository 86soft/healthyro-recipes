package app

import (
	"github.com/86soft/healthyro-recipes/app/query"
	"github.com/86soft/healthyro-recipes/domain"
)

type IdentifiableQuery interface {
	GetQueryIDPayload() string
}

type Queries struct {
	GetRecipeById query.GetRecipeByIdHandler
	ListRecipes   query.ListRecipesHandler
}

func NewQueryHandlers(repo domain.Repository) Queries {
	return Queries{
		GetRecipeById: query.NewGetRecipeByIdHandler(repo),
		ListRecipes:   query.NewListRecipesHandler(repo),
	}
}
