package engine

// Handler ...
type Handler func(c *Context) error

// HandlerRegistration ...
type HandlerRegistration struct {
	Service string // service name of handler
	Enabled bool
	Handler Handler
}

// HandlerRegistry ...
type HandlerRegistry interface {
	Handlers() []*HandlerRegistration
}
