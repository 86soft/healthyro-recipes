package domain

import (
	"fmt"
	"github.com/86soft/healthyro-recipes/app/command"
	"github.com/86soft/healthyro-recipes/app/query"
	"github.com/google/uuid"
)

type RID struct {
	id uuid.UUID
}

// NilRID is only for returning values on error
var NilRID = RID{id: uuid.Nil}

func NewRID(id string) (RID, error) {
	rid, err := uuid.Parse(id)
	if err != nil {
		return NilRID, fmt.Errorf("invalid uuid: %s", id)
	}
	return RID{id: rid}, nil
}

func NewRIDFromCmd(cmd command.IdentifiableCommand) (RID, error) {
	id := cmd.GetCommandIDPayload()
	return NewRID(id)
}

func NewRIDFromQuery(q query.IdentifiableQuery) (RID, error) {
	id := q.GetQueryIDPayload()
	return NewRID(id)
}

func (rid RID) GetID() uuid.UUID {
	return rid.id
}
