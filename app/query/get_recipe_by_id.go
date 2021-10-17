package query

import "github.com/google/uuid"

type GetRecipeById struct {
	RecipeUUID uuid.UUID
}
