package domain

import "github.com/google/uuid"

type ResourceID struct {
}

type RecipeID interface {
}

type TagID struct {
}

type ID[T any] struct {
	ID string
}

func CreateID[T any]() ID[T] {
	return ID[T]{ID: uuid.New().String()}
}

func FromStringID[T any](s string) ID[T] {
	return ID[T]{ID: s}
}
