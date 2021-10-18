package recipe

import "context"

type Repository interface {
	AddRecipe
	GetRecipe
	GetRecipes
	UpdateRecipe
	DeleteRecipe
}

type AddRecipe interface {
	AddRecipe(ctx context.Context, title string, description string, externalLink string) error
}
type GetRecipe interface {
	GetRecipe(ctx context.Context, recipeUUID string) (Recipe, error)
}
type GetRecipes interface {
	GetRecipes(ctx context.Context) ([]Recipe, error)
}
type UpdateRecipe interface {
	UpdateRecipe(ctx context.Context, r Recipe) error
}
type DeleteRecipe interface {
	DeleteRecipe(ctx context.Context, recipeUUID string) error
}
