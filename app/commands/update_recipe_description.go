package commands

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	l "github.com/rs/zerolog"
)

type UpdateRecipeDescription struct {
	RecipeID    d.ID[d.Recipe]
	Description string
}

type UpdateRecipeDescriptionHandler struct {
	updateDescriptionFn func(ctx context.Context, id d.ID[d.Recipe], description string) error
	logger              l.Logger
}

func NewUpdateRecipeDescriptionHandler(
	fn func(
		ctx context.Context,
		id d.ID[d.Recipe],
		description string,
	) error,
	logger l.Logger) (UpdateRecipeDescriptionHandler, error) {
	if fn == nil {
		return UpdateRecipeDescriptionHandler{}, &d.NilDependencyError{
			Name: "UpdateRecipeDescriptionHandler - fn",
		}
	}
	return UpdateRecipeDescriptionHandler{updateDescriptionFn: fn, logger: logger}, nil
}

func (h *UpdateRecipeDescriptionHandler) Handle(ctx context.Context, cmd UpdateRecipeDescription) error {
	if len(cmd.Description) > d.DescriptionLengthLimit {
		return d.ErrLengthLimitExceeded
	}
	return h.updateDescriptionFn(ctx, cmd.RecipeID, cmd.Description)
}
