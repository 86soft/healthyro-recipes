package app

import (
	"github.com/86soft/healthyro-recipes/app/queries"
	"github.com/86soft/healthyro-recipes/domain"
	"github.com/rs/zerolog"
)

type Queries struct {
	GetRecipeById queries.GetRecipeByIdHandler
	ListRecipes   queries.ListRecipesHandler
}

func NewQueryHandlers(repo domain.Store, logger zerolog.Logger) (Queries, error) {
	// t, _ := queries.NewGetRecipeByIdHandler(repo)
	return Queries{}, nil
}
