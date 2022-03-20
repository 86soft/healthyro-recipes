package core

import (
	"fmt"
)

type NilDependencyError struct {
	Name string
}

func (e *NilDependencyError) Error() string {
	return fmt.Sprintf("missing dependency inside %s", e.Name)
}

type CorruptedUUIDError struct {
	ID      string
	Details string
}

func (e *CorruptedUUIDError) Error() string {
	return fmt.Sprintf("corrupted uuid: %s: %s", e.ID, e.Details)
}

type OnDBUpdateError struct {
	ID      string
	Details string
}

func (e *OnDBUpdateError) Error() string {
	return fmt.Sprintf("could not update item with uuid %s: %s", e.ID, e.Details)
}
