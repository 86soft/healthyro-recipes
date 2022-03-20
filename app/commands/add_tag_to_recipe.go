package commands

import (
	"context"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type AddTagToRecipe struct {
	Name         string
	TagID        c.ID[c.Tag]
	RecipeID     c.ID[c.Recipe]
	CreateNewTag bool
}

type AddTagToRecipeHandler struct {
	addRecipeTagFn func(ctx context.Context, id c.ID[c.Recipe], t *c.Tag) error
	createTagFn    func(ctx context.Context, name string) (c.ID[c.Tag], error)
	log            zerolog.Logger
}

func NewAddTagToRecipeHandler(
	addTag func(ctx context.Context, id c.ID[c.Recipe], t *c.Tag) error,
	createTag func(ctx context.Context, name string) (c.ID[c.Tag], error),
	logger zerolog.Logger,
) (AddTagToRecipeHandler, error) {
	if addTag == nil {
		return AddTagToRecipeHandler{}, &c.NilDependencyError{
			Name: "RemoveTagFromRecipeHandler - addTag",
		}
	}
	if createTag == nil {
		return AddTagToRecipeHandler{}, &c.NilDependencyError{
			Name: "RemoveTagFromRecipeHandler - createTag",
		}
	}
	return AddTagToRecipeHandler{
		addRecipeTagFn: addTag,
		createTagFn:    createTag,
		log:            logger,
	}, nil
}

func (h *AddTagToRecipeHandler) Handle(ctx context.Context, cmd AddTagToRecipe) error {
	tag, err := h.GetOrCreateTag(ctx, cmd)
	if err != nil {
		return err
	}
	return h.addRecipeTagFn(ctx, cmd.RecipeID, &tag)
}

func (h *AddTagToRecipeHandler) GetOrCreateTag(ctx context.Context, cmd AddTagToRecipe) (c.Tag, error) {
	var err error
	tagID := cmd.TagID
	if cmd.CreateNewTag {
		tagID, err = h.createTagFn(ctx, cmd.Name)
	}
	return c.Tag{
		ID:       tagID,
		RecipeId: cmd.RecipeID,
		Name:     cmd.Name,
	}, err
}
