package unit

import (
	"net/http"

	"github.com/bitwormhole/tbmp/engine"
)

// DemoHandler ... 示例 handler
type DemoHandler struct {

	//starter:component
	_as func(engine.HandlerRegistry) //starter:as(".")

}

func (inst *DemoHandler) _impl() engine.HandlerRegistry {
	return inst
}

// Handlers ...
func (inst *DemoHandler) Handlers() []*engine.HandlerRegistration {

	r1 := &engine.HandlerRegistration{
		Method:  http.MethodGet,
		Path:    "/demo/1",
		Enabled: true,
		Handler: inst.handle,
	}

	return []*engine.HandlerRegistration{r1}
}

// Units ...
func (inst *DemoHandler) handle(ctx *engine.Context) error {
	return nil
}
