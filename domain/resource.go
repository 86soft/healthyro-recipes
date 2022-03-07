package domain

type Resource struct {
	id       ResourceID
	recipeID RecipeID
	name     string
	kind     string
	value    string
}
