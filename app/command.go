package app

import (
	"github.com/86soft/healthyro-recipes/app/commands"
	"github.com/86soft/healthyro-recipes/domain"
)

type CommandHandlers struct {
	CreateRecipe            commands.CreateRecipeHandler
	UpdateRecipeTitle       commands.UpdateRecipeTitleHandler
	UpdateRecipeDescription commands.UpdateRecipeDescriptionHandler
	DeleteRecipe            commands.DeleteRecipeHandler
}

func NewCommandHandlers(repo domain.Repository) CommandHandlers {
	return CommandHandlers{
		CreateRecipe:            commands.NewCreateRecipeHandler(repo),
		UpdateRecipeTitle:       commands.NewUpdateRecipeTitleHandler(repo),
		UpdateRecipeDescription: commands.NewUpdateRecipeDescriptionHandler(repo),
		DeleteRecipe:            commands.NewDeleteRecipeHandler(repo),
	}
}
