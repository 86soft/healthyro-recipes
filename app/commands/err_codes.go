package commands

import (
	"fmt"
)

type NilDependencyError struct {
	name string
}

func (e *NilDependencyError) Error() string {
	return fmt.Sprintf("missing dependency inside %s", e.name)
}
