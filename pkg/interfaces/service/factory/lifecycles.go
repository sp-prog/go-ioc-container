package factory

type Lifecycle int

const (
	Transient Lifecycle = iota
	Singleton
	Scoped
)
