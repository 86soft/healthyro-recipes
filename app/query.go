package app

import (
	"github.com/86soft/healthyro-recipes/app/queries"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type Queries struct {
	GetRecipeById            queries.GetRecipeByIdHandler
	ListRecipes              queries.ListRecipesHandler
	FindRecipesByName        queries.FindRecipesByNameHandler
	FindRecipesByTags        queries.FindRecipesByTagsHandler
	FindRecipesByNameAndTags queries.FindRecipesByNameAndTagsHandler
}

func NewQueryHandlers(repo core.Store, logger zerolog.Logger) (Queries, error) {
	getRecipe, err := queries.NewGetRecipeByIdHandler(repo.GetRecipe, logger)
	if err != nil {
		return Queries{}, err
	}
	listRecipes, err := queries.NewListRecipesHandler(repo.ListRecipes, logger)
	if err != nil {
		return Queries{}, err
	}
	findRecipesByName, err := queries.NewFindRecipesByNameHandler(repo.FindRecipesByName, logger)
	if err != nil {
		return Queries{}, err
	}
	findRecipesByTags, err := queries.NewFindRecipesByTagsHandler(repo.FindRecipesByTags, logger)
	if err != nil {
		return Queries{}, err
	}
	findRecipesByNameAndTags, err := queries.NewFindRecipesByNameAndTagsHandler(repo.FindRecipesByNameAndTags, logger)
	if err != nil {
		return Queries{}, err
	}

	return Queries{
		GetRecipeById:            getRecipe,
		ListRecipes:              listRecipes,
		FindRecipesByName:        findRecipesByName,
		FindRecipesByTags:        findRecipesByTags,
		FindRecipesByNameAndTags: findRecipesByNameAndTags,
	}, nil
}
