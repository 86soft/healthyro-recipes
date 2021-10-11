package command

import uuid "github.com/satori/go.uuid"

type UpdateRecipeDescription struct {
	id          uuid.UUID
	description string
}

type UpdateRecipeDescriptionHandler struct {
}
