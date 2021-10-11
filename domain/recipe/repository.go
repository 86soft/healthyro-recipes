package recipe

import "context"

type Repository interface {
	AddRecipe
	GetRecipe
}

type AddRecipe interface {
	AddRecipe(ctx context.Context, r *Recipe) error
}
type GetRecipe interface {
	GetRecipe(ctx context.Context) (*Recipe, error)
}
