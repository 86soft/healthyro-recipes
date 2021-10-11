package command

import uuid "github.com/satori/go.uuid"

type UpdateRecipeExternalLink struct {
	RecipeUUID   uuid.UUID
	ExternalLink string
}

type UpdateRecipeExternalLinkHandler struct {
}
