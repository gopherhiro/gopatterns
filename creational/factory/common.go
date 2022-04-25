package factory

type Course interface {
	GetID() int64
	GetName() string
}
