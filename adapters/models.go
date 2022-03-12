package adapters

import (
	"time"
)

type Document struct {
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
	DeletedAt time.Time `bson:"deletedAt,omitempty"`
}

type Recipe struct {
	Document    `bson:",inline"`
	Id          string      `bson:"_id"`
	Title       string      `bson:"title"`
	Description string      `bson:"description,omitempty"`
	Resources   []Resource  `bson:"resources,omitempty"`
	Tags        []RecipeTag `bson:"tags,omitempty"`
}
type RecipeTag struct {
	Id   string `bson:"_id"`
	name string
}

type Resource struct {
	Document `bson:",inline"`
	Id       string `bson:"_id"`
	Name     string
	Kind     string `bson:"kind,omitempty"`
	Value    string `bson:"value,omitempty"`
}

type Tag struct {
	Document  `bson:",inline"`
	ID        string   `bson:"_id"`
	Name      string   `bson:"name"`
	RecipeIDS []string `bson:"recipeIds"`
}
