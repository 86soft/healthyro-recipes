package command

type IdentifiableCommand interface {
	GetCommandIDPayload() string
}
