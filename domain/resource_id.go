package domain

type ResourceID struct {
	Id string
}

func (i ResourceID) GetID() string {
	return i.Id
}
