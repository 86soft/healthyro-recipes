package core

import (
	"fmt"
	"github.com/segmentio/ksuid"
)

type ID[T any] struct {
	Value ksuid.KSUID
}

func CreateID[T any]() ID[T] {
	return ID[T]{Value: ksuid.New()}
}

func FromStringID[T any](s string) (ID[T], error) {
	id, err := ksuid.Parse(s)
	if err != nil {
		return ID[T]{}, fmt.Errorf("FromStringID: %w", err)
	}
	return ID[T]{Value: id}, nil
}
