package app

import (
	"github.com/86soft/healthyro-recipes/app/commands"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type CommandHandlers struct {
	AddRecipe               commands.AddRecipeHandler
	UpdateRecipeTitle       commands.UpdateRecipeTitleHandler
	UpdateRecipeDescription commands.UpdateRecipeDescriptionHandler
	DeleteRecipe            commands.DeleteRecipeHandler
}

func NewCommandHandlers(repo core.Store, logger zerolog.Logger) (CommandHandlers, error) {
	arh, err := commands.NewAddRecipeHandler(repo, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	//crh, err := commands.NewUpdateRecipeTitleHandler(repo, logger)
	if err != nil {
		return CommandHandlers{}, err
	}

	return CommandHandlers{
		AddRecipe: arh,
		//UpdateRecipeTitle:       commands.NewUpdateRecipeTitleHandler(repo),
		//UpdateRecipeDescription: commands.NewUpdateRecipeDescriptionHandler(repo),
		//DeleteRecipe:            commands.NewDeleteRecipeHandler(repo),
	}, nil
}
