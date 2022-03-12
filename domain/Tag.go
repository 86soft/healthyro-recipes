package domain

type TagID struct {
	id string
}

func NewTagID(id string) TagID {
	return TagID{id: id}
}

func (t *TagID) GetID() string {
	return t.id
}

type Tag struct {
	Id       TagID
	RecipeId RecipeID
	Name     string
}
