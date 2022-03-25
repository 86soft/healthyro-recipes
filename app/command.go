package app

import (
	"github.com/86soft/healthyro-recipes/app/commands"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type CommandHandlers struct {
	CreateRecipe             commands.CreateRecipeHandler
	AddRecipeResource        commands.AddRecipeResourceHandler
	AddTagToRecipe           commands.AddTagToRecipeHandler
	CreateTag                commands.CreateTagHandler
	DeleteRecipe             commands.DeleteRecipeHandler
	RemoveResourceFromRecipe commands.RemoveResourceFromRecipeHandler
	RemoveTagFromRecipe      commands.RemoveTagFromRecipeHandler
	UpdateRecipeTitle        commands.UpdateRecipeTitleHandler
	UpdateRecipeDescription  commands.UpdateRecipeDescriptionHandler
}

func NewCommandHandlers(repo core.Store, logger zerolog.Logger) (CommandHandlers, error) {
	createRecipe, err := commands.NewCreateRecipeHandler(repo.CreateRecipe, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	addRecipeResource, err := commands.NewAddRecipeResourceHandler(repo.AddRecipeResource, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	addTagToRecipe, err := commands.NewAddTagToRecipeHandler(repo.AddTagToRecipe, repo.CreateTag, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	createTag, err := commands.NewCreateTagHandler(repo.CreateTag, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	removeResourceFromRecipe, err := commands.NewRemoveResourceFromRecipeHandler(repo.RemoveResourceFromRecipe, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	removeTagFromRecipe, err := commands.NewRemoveTagFromRecipeHandler(repo.RemoveTagFromRecipe, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	updateRecipeTitle, err := commands.NewUpdateRecipeTitleHandler(repo.UpdateRecipeTitle, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	updateRecipeDescription, err := commands.NewUpdateRecipeDescriptionHandler(repo.UpdateRecipeDescription, logger)
	if err != nil {
		return CommandHandlers{}, err
	}
	deleteRecipe, err := commands.NewDeleteRecipeHandler(repo.DeleteRecipe, logger)
	if err != nil {
		return CommandHandlers{}, err
	}

	return CommandHandlers{
		CreateRecipe:             createRecipe,
		AddRecipeResource:        addRecipeResource,
		AddTagToRecipe:           addTagToRecipe,
		CreateTag:                createTag,
		DeleteRecipe:             deleteRecipe,
		RemoveResourceFromRecipe: removeResourceFromRecipe,
		RemoveTagFromRecipe:      removeTagFromRecipe,
		UpdateRecipeTitle:        updateRecipeTitle,
		UpdateRecipeDescription:  updateRecipeDescription,
	}, nil
}
