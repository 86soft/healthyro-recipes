package command

import "github.com/google/uuid"

type UpdateRecipeExternalLink struct {
	RecipeUUID   uuid.UUID
	ExternalLink string
}

type UpdateRecipeExternalLinkHandler struct {
}
