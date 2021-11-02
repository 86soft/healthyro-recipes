package app

import (
	"github.com/86soft/healthyro-recipes/app/command"
	"github.com/86soft/healthyro-recipes/domain"
)

type IdentifiableCommand interface {
	GetCommandIDPayload() string
}

type Commands struct {
	CreateRecipe             command.CreateRecipeHandler
	UpdateRecipeTitle        command.UpdateRecipeTitleHandler
	UpdateRecipeDescription  command.UpdateRecipeDescriptionHandler
	UpdateRecipeExternalLink command.UpdateRecipeExternalLinkHandler
	DeleteRecipe             command.DeleteRecipeHandler
}

func NewCommandHandlers(repo domain.Repository) Commands {
	return Commands{
		CreateRecipe:             command.NewCreateRecipeHandler(repo),
		UpdateRecipeTitle:        command.NewUpdateRecipeTitleHandler(repo),
		UpdateRecipeDescription:  command.NewUpdateRecipeDescriptionHandler(repo),
		UpdateRecipeExternalLink: command.NewUpdateRecipeExternalLinkHandler(repo),
		DeleteRecipe:             command.NewDeleteRecipeHandler(repo),
	}
}
