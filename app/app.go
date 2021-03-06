package app

import (
	"fmt"
	"github.com/86soft/healthyro-recipes/core"
	"github.com/rs/zerolog"
)

type Application struct {
	Commands
	Queries
	Log zerolog.Logger
}

func NewApplication(repo core.Store, logger zerolog.Logger) (Application, error) {
	c, err := NewCommandHandlers(repo, logger)
	if err != nil {
		return Application{}, fmt.Errorf("NewApplication: %w", err)
	}
	q, err := NewQueryHandlers(repo, logger)
	if err != nil {
		return Application{}, fmt.Errorf("NewApplication: %w", err)
	}

	return Application{Queries: q, Commands: c, Log: logger}, nil
}
