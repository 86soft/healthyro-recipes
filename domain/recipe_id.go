package domain

type RecipeID struct {
	id string
}

func NewRecipeID(id string) RecipeID {
	return RecipeID{id}
}

func (rid RecipeID) GetID() string {
	return rid.id
}
