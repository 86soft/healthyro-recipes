package app

import (
	"github.com/86soft/healthyro-recipes/domain"
)

type Application struct {
	Commands CommandHandlers
	Queries  Queries
}

func NewApplication(repo domain.Repository) Application {
	return Application{
		Commands: NewCommandHandlers(repo),
		Queries:  NewQueryHandlers(repo),
	}
}
