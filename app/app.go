package app

import (
	"fmt"
	"github.com/86soft/healthyro-recipes/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	Commands CommandHandlers
	Queries  Queries
	Log      zerolog.Logger
}

func NewApplication(repo domain.Repository, logger zerolog.Logger) (Application, error) {
	c, err := NewCommandHandlers(repo, logger)
	if err != nil {
		return Application{}, fmt.Errorf("NewApplication: %w", err)
	}

	q, err := NewQueryHandlers(repo, logger)
	if err != nil {
		return Application{}, fmt.Errorf("NewApplication: %w", err)
	}

	return Application{
		Commands: c,
		Queries:  q,
		Log:      logger,
	}, nil
}
