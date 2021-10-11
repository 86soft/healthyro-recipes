package command

import uuid "github.com/satori/go.uuid"

type DeleteRecipe struct {
	id uuid.UUID
}

type DeleteRecipeHandler struct {
}
