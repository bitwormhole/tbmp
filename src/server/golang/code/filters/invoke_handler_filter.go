package filters

import (
	"fmt"

	"github.com/bitwormhole/tbmp/engine"
)

// InvokeHandlerFilter ...
type InvokeHandlerFilter struct {

	//starter:component
	_as func(engine.FilterRegistry) //starter:as(".")

}

func (inst *InvokeHandlerFilter) _impl() (engine.FilterRegistry, engine.Filter) {
	return inst, inst
}

// Filters ...
func (inst *InvokeHandlerFilter) Filters() []*engine.FilterRegistration {
	r1 := &engine.FilterRegistration{
		Enabled: true,
		Order:   engine.FilterOrderInvokeHandler,
		Filter:  inst,
	}
	return []*engine.FilterRegistration{r1}
}

// DoFilter ...
func (inst *InvokeHandlerFilter) DoFilter(ctx *engine.Context, next engine.FilterChain) error {

	sel := ctx.Selector
	h := ctx.Handler
	if h == nil {
		return fmt.Errorf("tbmp: no handler for selector [%s]", sel)
	}

	err := h(ctx)
	if err != nil {
		return err
	}

	return next.DoFilter(ctx)
}
