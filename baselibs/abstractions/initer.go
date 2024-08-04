package abstractions

type Initer interface {
	// Init
	// call it before you use this instance.
	Init() error
}
