package command

import uuid "github.com/satori/go.uuid"

type UpdateRecipeDescription struct {
	RecipeUUID  uuid.UUID
	Description string
}

type UpdateRecipeDescriptionHandler struct {
}
