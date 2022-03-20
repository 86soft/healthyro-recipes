package commands

import (
	"context"
	c "github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type UpdateRecipeTitle struct {
	RecipeID c.ID[c.Recipe]
	Title    string
}

type UpdateRecipeTitleHandler struct {
	updateRecipeFn func(ctx context.Context, id c.ID[c.Recipe], title string) error
	logger         zerolog.Logger
}

func NewUpdateRecipeTitleHandler(
	fn func(ctx context.Context, id c.ID[c.Recipe], title string) error,
	logger zerolog.Logger,
) UpdateRecipeTitleHandler {
	if fn == nil {
		panic("nil updateDescriptionFn inside NewUpdateRecipeTitleHandler")
	}

	return UpdateRecipeTitleHandler{
		updateRecipeFn: fn,
		logger:         logger,
	}
}

func (h *UpdateRecipeTitleHandler) Handle(ctx context.Context, cmd UpdateRecipeTitle) error {
	if len(cmd.Title) > c.TitleLengthLimit {
		return c.ErrLengthLimitExceeded
	}
	return h.updateRecipeFn(ctx, cmd.RecipeID, cmd.Title)
}
