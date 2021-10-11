package app

import "github.com/86soft/healthyro-recipes/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateRecipe             command.CreateRecipe
	UpdateRecipeTitle        command.UpdateRecipeTitle
	UpdateRecipeDescription  command.UpdateRecipeDescription
	UpdateRecipeExternalLink command.UpdateRecipeExternalLink
	DeleteRecipe             command.DeleteRecipe
}

type Queries struct {
}
