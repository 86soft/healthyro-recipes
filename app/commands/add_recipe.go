package commands

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
	uuid "github.com/google/uuid"
	"github.com/rs/zerolog"
)

type AddRecipe struct {
	Title       string
	Description string
	Resources   []RecipeResources
	Tags        []string
}

type RecipeResources struct {
	Name  string
	Kind  string
	Value string
}

type AddRecipeHandler struct {
	addRecipeFn func(ctx context.Context, newRecipe *domain.Recipe) error
	log         zerolog.Logger
}

func NewAddRecipeHandler(db domain.Repository, logger zerolog.Logger) (AddRecipeHandler, error) {
	if db == nil {
		return AddRecipeHandler{}, &NilDependencyError{
			name: "NewAddRecipeHandler - db",
		}
	}
	return AddRecipeHandler{
		addRecipeFn: db.AddRecipe,
		log:         logger,
	}, nil
}

func (h *AddRecipeHandler) Handle(ctx context.Context, cmd AddRecipe) (domain.RecipeID, error) {
	id := domain.NewRecipeID(uuid.New().String())
	resources := make([]domain.Resource, 0, len(cmd.Resources))
	for _, r := range cmd.Resources {
		resources = append(resources, domain.Resource{
			Id:    domain.ResourceID{Id: uuid.New().String()},
			Name:  r.Name,
			Kind:  r.Kind,
			Value: r.Value,
		})
	}

	tags := make([]domain.Tag, 0, len(cmd.Tags))
	for _, t := range cmd.Resources {
		tags = append(tags, domain.Tag{
			Id:   domain.NewTagID(uuid.New().String()),
			Name: t.Name,
		})
	}

	recipe := domain.Recipe{
		Id:          id,
		Title:       cmd.Title,
		Description: cmd.Description,
		Resources:   resources,
		Tags:        tags,
	}
	err := h.addRecipeFn(ctx, &recipe)
	if err != nil {
		return domain.RecipeID{}, err
	}
	return id, nil
}
