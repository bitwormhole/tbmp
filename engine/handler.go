package engine

import "github.com/bitwormhole/tbmp"

// Handler ...
type Handler func(c *Context) error

////////////////////////////////////////////////////////////////////////////////

// HandlerRegistration ...
type HandlerRegistration struct {
	Method  string // for selector
	Path    string // for selector
	Enabled bool
	Handler Handler
}

// GetSelector ...
func (inst *HandlerRegistration) GetSelector() tbmp.HandlerSelector {
	b := &tbmp.HandlerSelectorBuilder{
		Method: inst.Method,
		Path:   inst.Path,
	}
	return b.Create()
}

////////////////////////////////////////////////////////////////////////////////

// HandlerRegistry ...
type HandlerRegistry interface {
	Handlers() []*HandlerRegistration
}
