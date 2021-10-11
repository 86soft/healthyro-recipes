package command

import uuid "github.com/satori/go.uuid"

type UpdateRecipeTitle struct {
	RecipeUUID uuid.UUID
	Title      string
}

type UpdateRecipeTitleHandler struct {
}
