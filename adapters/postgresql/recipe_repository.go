package postgresql

import (
	"context"
	domain "github.com/86soft/healthyro-recipes/domain/recipe"
	"gorm.io/gorm"
)

type RecipeRepository struct {
	db gorm.DB
}

func (repo *RecipeRepository) AddRecipe(ctx context.Context,
	title string, description string, externalLink string) error {
	return nil
}

func (repo *RecipeRepository) GetRecipe(ctx context.Context) (*domain.Recipe, error) {
	return nil, nil
}

func (repo *RecipeRepository) GetRecipes(ctx context.Context) (*domain.Recipe, error) {
	return nil, nil
}
