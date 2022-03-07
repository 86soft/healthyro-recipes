package domain

import (
	"context"
)

type Repository interface {
	AddRecipe(ctx context.Context, r *Recipe) error
	GetRecipe(ctx context.Context, id RecipeID) (Recipe, error)
	GetRecipes(ctx context.Context) ([]Recipe, error)
	UpdateRecipeTitle(ctx context.Context, id RecipeID, title string) error
	UpdateRecipeDescription(ctx context.Context, id RecipeID, description string) error
	DeleteRecipe(ctx context.Context, id RecipeID) error
	AddRecipeResource(ctx context.Context, id RecipeID, r *Resource) error
	DeleteRecipeResource(ctx context.Context, id RecipeID, name string) error
}
