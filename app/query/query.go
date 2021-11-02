package query

type IdentifiableQuery interface {
	GetQueryIDPayload() string
}
