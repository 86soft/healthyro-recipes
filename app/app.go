package app

import (
	"github.com/86soft/healthyro-recipes/app/command"
	"github.com/86soft/healthyro-recipes/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateRecipe             command.CreateRecipeHandler
	UpdateRecipeTitle        command.UpdateRecipeTitleHandler
	UpdateRecipeDescription  command.UpdateRecipeDescriptionHandler
	UpdateRecipeExternalLink command.UpdateRecipeExternalLinkHandler
	DeleteRecipe             command.DeleteRecipeHandler
}

type Queries struct {
	GetRecipeById query.GetRecipeByIdHandler
	ListRecipes   query.ListRecipesHandler
}
