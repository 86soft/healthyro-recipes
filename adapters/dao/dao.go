package dao

import (
	"github.com/86soft/healthyro-recipes/domain"
)

type RecipeModel struct {
	Base
	title        string
	description  string
	externalLink string
}

func NewModelWithID(rid domain.RID) RecipeModel {
	return RecipeModel{
		Base: Base{
			ID: rid.GetID(),
		},
	}
}

func MapRecipeToModel(rcp *domain.Recipe) RecipeModel {
	return RecipeModel{
		Base: Base{
			ID: rcp.RecipeID(),
		},
		title:        rcp.Title(),
		description:  rcp.Description(),
		externalLink: rcp.ExternalLink(),
	}
}

func MapModelToRecipe(dao *RecipeModel) domain.Recipe {
	// we can ignore error because we can't store nil uuid in db
	id, _ := domain.NewRIDFromUUID(dao.ID)
	return domain.UnmarshalRecipe(id, dao.title, dao.description, dao.externalLink)

}

func MapModelsToRecipe(dao []RecipeModel) (res []domain.Recipe) {
	for _, model := range dao {
		res = append(res, MapModelToRecipe(&model))
	}
	return res
}
