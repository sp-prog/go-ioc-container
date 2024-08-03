package type_factory

type Lifecycle int

const (
	Transient Lifecycle = iota
	Singleton
	Scoped
)
