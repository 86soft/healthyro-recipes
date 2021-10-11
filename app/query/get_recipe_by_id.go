package query

import uuid "github.com/satori/go.uuid"

type GetRecipeById struct {
	RecipeUUID uuid.UUID
}
