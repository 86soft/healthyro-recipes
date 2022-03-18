package domain

type Resource struct {
	ID    ID[Resource]
	Name  string
	Kind  string
	Value string
}
