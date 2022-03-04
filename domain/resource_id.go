package domain

type ResourceID struct {
	id string
}

func (i ResourceID) GetID() string {
	return i.id
}
