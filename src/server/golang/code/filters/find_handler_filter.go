package filters

import (
	"fmt"
	"sync"

	"github.com/bitwormhole/tbmp"
	"github.com/bitwormhole/tbmp/engine"
	"github.com/starter-go/vlog"
)

// FindHandlerFilter ...
type FindHandlerFilter struct {

	//starter:component
	_as func(engine.FilterRegistry, engine.ElementRegistry) //starter:as(".",".")

	mutex    sync.Mutex
	handlers map[tbmp.HandlerSelector]*engine.HandlerRegistration
}

func (inst *FindHandlerFilter) _impl() (engine.FilterRegistry, engine.ElementRegistry, engine.Filter) {
	return inst, inst, inst
}

// Elements ...
func (inst *FindHandlerFilter) Elements() []*engine.ElementRegistration {
	r1 := &engine.ElementRegistration{
		Enabled: true,
		OnInit:  inst.init,
	}
	return []*engine.ElementRegistration{r1}
}

// Filters ...
func (inst *FindHandlerFilter) Filters() []*engine.FilterRegistration {
	r1 := &engine.FilterRegistration{
		Enabled: true,
		Order:   engine.FilterOrderFindHandler,
		Filter:  inst,
	}
	return []*engine.FilterRegistration{r1}
}

func (inst *FindHandlerFilter) init(eng *engine.Engine) error {
	src := eng.Handlers
	dst := make(map[tbmp.HandlerSelector]*engine.HandlerRegistration)
	for _, item := range src {
		if inst.isHandlerReady(item) {
			sel := item.GetSelector()
			old := dst[sel]
			if old != nil {
				vlog.Warn("tbmp: handlers with selector [%s] are duplicate", sel)
			}
			dst[sel] = item
		}
	}
	inst.handlers = dst
	return nil
}

func (inst *FindHandlerFilter) isHandlerReady(hr *engine.HandlerRegistration) bool {
	if hr == nil {
		return false
	}
	if !hr.Enabled {
		return false
	}
	if hr.Handler == nil {
		return false
	}
	// if hr.Service == "" {
	// 	return false
	// }
	return true
}

// DoFilter ...
func (inst *FindHandlerFilter) DoFilter(c *engine.Context, next engine.FilterChain) error {

	h, err := inst.findHandler(c)
	if err != nil {
		return err
	}
	c.Handler = h

	return next.DoFilter(c)
}

func (inst *FindHandlerFilter) findHandler(c *engine.Context) (engine.Handler, error) {

	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	sel := c.Selector
	h1 := inst.handlers[sel]
	if h1 == nil {
		err := inst.makeNoHandlerError(sel)
		return nil, err
	}

	h2 := h1.Handler
	if h2 == nil {
		err := inst.makeNoHandlerError(sel)
		return nil, err
	}

	return h2, nil
}

func (inst *FindHandlerFilter) makeNoHandlerError(sel tbmp.HandlerSelector) error {
	return fmt.Errorf("no handler for service [%s]", sel)
}
