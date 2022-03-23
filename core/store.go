package core

import (
	"context"
)

type Store interface {
	RecipeStore
	TagStore
	ResourceStore
}
type RecipeStore interface {
	ListRecipes(ctx context.Context) ([]Recipe, error)
	FindRecipesByName(ctx context.Context, name string) ([]Recipe, error)
	GetRecipe(ctx context.Context, id ID[Recipe]) (Recipe, error)
	CreateRecipe(ctx context.Context, r *Recipe) error
	UpdateRecipeTitle(ctx context.Context, id ID[Recipe], title string) error
	UpdateRecipeDescription(ctx context.Context, id ID[Recipe], description string) error
	DeleteRecipe(ctx context.Context, id ID[Recipe]) error
}

type ResourceStore interface {
	DeleteRecipeResource(ctx context.Context, recipeID ID[Recipe], resourceID ID[Resource]) error
	AddRecipeResource(ctx context.Context, id ID[Recipe], r *Resource) error
}

type TagStore interface {
	CreateTag(ctx context.Context, name string) (ID[Tag], error)
	AddTagToRecipe(ctx context.Context, id ID[Recipe], t *Tag) error
	RemoveTagFromRecipe(ctx context.Context, recipeID ID[Recipe], tagID ID[Tag]) error
}
