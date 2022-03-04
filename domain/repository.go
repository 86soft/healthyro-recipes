package domain

import (
	"context"
)

type Repository interface {
	AddRecipe
	GetRecipe
	GetRecipes
	UpdateRecipeTitle
	UpdateRecipeDescription
	UpdateRecipeExternalLink
	DeleteRecipe
}

type AddRecipe interface {
	AddRecipe(ctx context.Context, newRecipe *Recipe) (RecipeID, error)
}
type GetRecipe interface {
	GetRecipe(ctx context.Context, recipeID RecipeID) (Recipe, error)
}
type GetRecipes interface {
	GetRecipes(ctx context.Context) ([]Recipe, error)
}
type UpdateRecipeTitle interface {
	UpdateRecipeTitle(ctx context.Context, recipeID RecipeID, title string) error
}
type UpdateRecipeDescription interface {
	UpdateRecipeDescription(ctx context.Context, recipeID RecipeID, description string) error
}
type UpdateRecipeExternalLink interface {
	UpdateRecipeExternalLink(ctx context.Context, recipeID RecipeID, link string) error
}
type DeleteRecipe interface {
	DeleteRecipe(ctx context.Context, recipeID RecipeID) error
}
