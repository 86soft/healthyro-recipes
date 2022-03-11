package adapters

import (
	"time"
)

type Recipe struct {
	Id          string     `bson:"_id"`
	CreatedAt   time.Time  `bson:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at,omitempty"`
	DeletedAt   time.Time  `bson:"deleted_at,omitempty"`
	Title       string     `bson:"title,omitempty"`
	Description string     `bson:"description,omitempty"`
	Resources   []Resource `bson:"resources,omitempty"`
}

type Resource struct {
	Name  string `bson:"name"`
	Kind  string `bson:"kind,omitempty"`
	Value string `bson:"value,omitempty"`
}
