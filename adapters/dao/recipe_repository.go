package dao

import (
	"context"
	"github.com/86soft/healthyro-recipes/domain"
	"gorm.io/gorm"
)

type RecipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(conn *gorm.DB) RecipeRepository {
	return RecipeRepository{db: conn}
}

func (r RecipeRepository) AddRecipe(ctx context.Context, newRecipe *domain.Recipe) (domain.RID, error) {
	dao := MapRecipeToModel(newRecipe)
	err := r.db.Create(&dao).Error
	if err != nil {
		return domain.NilRID, err
	}
	return domain.NewRIDFromUUID(dao.ID)
}

func (r RecipeRepository) GetRecipe(ctx context.Context, recipeID domain.RID) (domain.Recipe, error) {
	dao := NewModelWithID(recipeID)
	err := r.db.Find(&dao).Error
	if err != nil {
		return domain.NilRecipe, err
	}
	return MapModelToRecipe(&dao), nil
}

func (r RecipeRepository) GetRecipes(ctx context.Context) ([]domain.Recipe, error) {
	var recipes []RecipeModel
	err := r.db.Model(RecipeModel{}).Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return MapModelsToRecipe(recipes), nil
}

func (r RecipeRepository) UpdateRecipeTitle(ctx context.Context, recipeID domain.RID, title string) error {
	changes := map[string]interface{}{"Title": title}
	dao := NewModelWithID(recipeID)
	return r.db.Model(&dao).Updates(changes).Error
}

func (r RecipeRepository) UpdateRecipeDescription(ctx context.Context, recipeID domain.RID, description string) error {
	changes := map[string]interface{}{"Description": description}
	dao := NewModelWithID(recipeID)
	return r.db.Model(&dao).Updates(changes).Error
}

func (r RecipeRepository) UpdateRecipeExternalLink(ctx context.Context, recipeID domain.RID, link string) error {
	changes := map[string]interface{}{"ExternalLink": link}
	dao := NewModelWithID(recipeID)
	return r.db.Model(&dao).Updates(changes).Error
}

func (r RecipeRepository) DeleteRecipe(ctx context.Context, recipeID domain.RID) error {
	dao := NewModelWithID(recipeID)
	return r.db.Delete(dao).Error
}
