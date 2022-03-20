package commands

import (
	"context"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type RemoveTagFromRecipe struct {
	TagID    c.ID[c.Tag]
	RecipeID c.ID[c.Recipe]
}

type RemoveTagFromRecipeHandler struct {
	removeTagFn func(ctx context.Context, recipeID c.ID[c.Recipe], tagID c.ID[c.Tag]) error
	log         zerolog.Logger
}

func NewRemoveTagFromRecipeHandler(
	removeTag func(ctx context.Context, recipeID c.ID[c.Recipe], tagID c.ID[c.Tag]) error,
	logger zerolog.Logger,
) (RemoveTagFromRecipeHandler, error) {
	if removeTag == nil {
		return RemoveTagFromRecipeHandler{}, &c.NilDependencyError{
			Name: "RemoveTagFromRecipeHandler - createTagFn",
		}
	}
	return RemoveTagFromRecipeHandler{
		removeTagFn: removeTag,
		log:         logger,
	}, nil
}

func (h *RemoveTagFromRecipeHandler) Handle(ctx context.Context, cmd AddTagToRecipe) error {
	return h.removeTagFn(ctx, cmd.RecipeID, cmd.TagID)
}
