package core

import (
	c "context"
)

type Store interface {
	RecipeStore
	TagStore
	ResourceStore
}

type (
	ListRecipes              func(ctx c.Context) ([]Recipe, error)
	FindRecipesByName        func(ctx c.Context, name string) ([]Recipe, error)
	FindRecipesByTags        func(ctx c.Context, tags []Tag) ([]Recipe, error)
	GetRecipe                func(ctx c.Context, id ID[Recipe]) (Recipe, error)
	CreateRecipe             func(ctx c.Context, r *Recipe) error
	UpdateRecipeTitle        func(ctx c.Context, id ID[Recipe], title string) error
	UpdateRecipeDescription  func(ctx c.Context, id ID[Recipe], description string) error
	DeleteRecipe             func(ctx c.Context, id ID[Recipe]) error
	RemoveResourceFromRecipe func(ctx c.Context, recipeID ID[Recipe], resourceID ID[Resource]) error
	AddRecipeResource        func(ctx c.Context, id ID[Recipe], r *Resource) error
	CreateTag                func(ctx c.Context, name string) (ID[Tag], error)
	AddTagToRecipe           func(ctx c.Context, id ID[Recipe], t *Tag) error
	RemoveTagFromRecipe      func(ctx c.Context, recipeID ID[Recipe], tagID ID[Tag]) error
)

type RecipeStore interface {
	ListRecipes(ctx c.Context) ([]Recipe, error)
	FindRecipesByName(ctx c.Context, name string) ([]Recipe, error)
	FindRecipesByTags(ctx c.Context, tags []Tag) ([]Recipe, error)
	GetRecipe(ctx c.Context, id ID[Recipe]) (Recipe, error)
	CreateRecipe(ctx c.Context, r *Recipe) error
	UpdateRecipeTitle(ctx c.Context, id ID[Recipe], title string) error
	UpdateRecipeDescription(ctx c.Context, id ID[Recipe], description string) error
	DeleteRecipe(ctx c.Context, id ID[Recipe]) error
}

type ResourceStore interface {
	RemoveResourceFromRecipe(ctx c.Context, recipeID ID[Recipe], resourceID ID[Resource]) error
	AddRecipeResource(ctx c.Context, id ID[Recipe], r *Resource) error
}

type TagStore interface {
	CreateTag(ctx c.Context, name string) (ID[Tag], error)
	AddTagToRecipe(ctx c.Context, id ID[Recipe], t *Tag) error
	RemoveTagFromRecipe(ctx c.Context, recipeID ID[Recipe], tagID ID[Tag]) error
}
