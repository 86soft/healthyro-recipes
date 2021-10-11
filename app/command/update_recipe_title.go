package command

import uuid "github.com/satori/go.uuid"

type UpdateRecipeTitle struct {
	id    uuid.UUID
	title string
}

type UpdateRecipeTitleHandler struct {
}
