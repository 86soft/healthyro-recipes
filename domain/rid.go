package domain

import (
	"errors"
	"fmt"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/google/uuid"
)

type RID struct {
	id uuid.UUID
}

// NilRID is only for returning values on error
var NilRID = RID{id: uuid.Nil}

var ErrNilUUID = errors.New("nil uuid")

func NewRIDFromString(id string) (RID, error) {
	newId, err := uuid.Parse(id)
	if err != nil {
		return NilRID, fmt.Errorf("invalid uuid: %s", id)
	}
	return NewRIDFromUUID(newId)
}

func NewRIDFromUUID(id uuid.UUID) (RID, error) {
	if id == uuid.Nil {
		return NilRID, ErrNilUUID
	}
	return RID{id: id}, nil
}

func NewRIDFromCmd(cmd app.IdentifiableCommand) (RID, error) {
	id := cmd.GetCommandIDPayload()
	return NewRIDFromString(id)
}

func NewRIDFromQuery(q app.IdentifiableQuery) (RID, error) {
	id := q.GetQueryIDPayload()
	return NewRIDFromString(id)
}

func (rid RID) GetID() uuid.UUID {
	return rid.id
}
