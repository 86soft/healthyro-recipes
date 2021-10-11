package command

import uuid "github.com/satori/go.uuid"

type DeleteRecipe struct {
	RecipeUUID uuid.UUID
}

type DeleteRecipeHandler struct {
}
