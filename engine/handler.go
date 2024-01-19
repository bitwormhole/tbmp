package engine

// Handler ...
type Handler func(c *Context)

// HandlerRegistration ...
type HandlerRegistration struct {
	Method  string
	Path    string
	Enabled bool
	Handler Handler
}

// HandlerRegistry ...
type HandlerRegistry interface {
	Handlers() []*HandlerRegistration
}
