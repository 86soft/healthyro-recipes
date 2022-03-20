package app

import (
	"github.com/86soft/healthyro-recipes/app/queries"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type Queries struct {
	GetRecipeById queries.GetRecipeByIdHandler
	ListRecipes   queries.ListRecipesHandler
}

func NewQueryHandlers(repo core.Store, logger zerolog.Logger) (Queries, error) {
	// t, _ := queries.NewGetRecipeByIdHandler(repo)
	return Queries{}, nil
}
