package command

import uuid "github.com/satori/go.uuid"

type UpdateRecipeExternalLink struct {
	id           uuid.UUID
	externalLink string
}

type UpdateRecipeExternalLinkHandler struct {
}
