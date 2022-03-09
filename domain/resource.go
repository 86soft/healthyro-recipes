package domain

type Resource struct {
	id       ResourceID
	recipeID RecipeID
	name     string
	kind     string
	value    string
}

func (r *Resource) Id() ResourceID {
	return r.id
}

func (r *Resource) RecipeID() RecipeID {
	return r.recipeID
}

func (r *Resource) Name() string {
	return r.name
}

func (r *Resource) Kind() string {
	return r.kind
}

func (r *Resource) Value() string {
	return r.value
}
