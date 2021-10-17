package command

import "github.com/google/uuid"

type DeleteRecipe struct {
	RecipeUUID uuid.UUID
}

type DeleteRecipeHandler struct {
}
