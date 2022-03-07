package adapters

import (
	"time"
)

type recipe struct {
	id          string     `bson:"_id"`
	createdAt   time.Time  `bson:"created_at"`
	updatedAt   time.Time  `bson:"updated_at,omitempty"`
	deletedAt   time.Time  `bson:"deleted_at,omitempty"`
	title       string     `bson:"title,omitempty"`
	description string     `bson:"description,omitempty"`
	resources   []resource `bson:"resources,omitempty"`
}

type resource struct {
	name  string `bson:"name,omitempty"`
	kind  string `bson:"kind,omitempty"`
	value string `bson:"value,omitempty"`
}
