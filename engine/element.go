package engine

// ElementOnInit ...
type ElementOnInit func(e *Engine) error

// ElementRegistration ...
type ElementRegistration struct {
	Enabled bool

	OnInit func(e *Engine) error
}

// ElementRegistry ...
type ElementRegistry interface {
	Elements() []*ElementRegistration
}
