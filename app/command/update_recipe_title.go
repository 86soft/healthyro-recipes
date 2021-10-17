package command

import "github.com/google/uuid"

type UpdateRecipeTitle struct {
	RecipeUUID uuid.UUID
	Title      string
}

type UpdateRecipeTitleHandler struct {
}
