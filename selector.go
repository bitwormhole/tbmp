package tbmp

import "strings"

// HandlerSelector 是一个字符串，用于选择 handler
type HandlerSelector string

////////////////////////////////////////////////////////////////////////////////

// HandlerSelectorBuilder ...
type HandlerSelectorBuilder struct {
	Method string
	Path   string
}

// Create ...
func (inst *HandlerSelectorBuilder) Create() HandlerSelector {
	mt := strings.ToUpper(inst.Method)
	str := mt + ":" + inst.Path
	return HandlerSelector(str)
}
