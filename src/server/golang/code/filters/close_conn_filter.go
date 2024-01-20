package filters

import (
	"github.com/bitwormhole/tbmp/engine"
)

// CloseConnectionFilter ...
type CloseConnectionFilter struct {

	//starter:component
	_as func(engine.FilterRegistry) //starter:as(".")

}

func (inst *CloseConnectionFilter) _impl() (engine.FilterRegistry, engine.Filter) {
	return inst, inst
}

// Filters ...
func (inst *CloseConnectionFilter) Filters() []*engine.FilterRegistration {
	r1 := &engine.FilterRegistration{
		Enabled: true,
		Order:   engine.FilterOrderCloseConnection,
		Filter:  inst,
	}
	return []*engine.FilterRegistration{r1}
}

// DoFilter ...
func (inst *CloseConnectionFilter) DoFilter(ctx *engine.Context, next engine.FilterChain) error {
	defer func() {
		inst.close(ctx)
	}()

	err := next.DoFilter(ctx)
	if err != nil {
		return err
	}

	down, err := ctx.Connection.Downstream()
	if err != nil {
		return err
	}

	down.Headers().Get("nop")
	return nil
}

func (inst *CloseConnectionFilter) close(ctx *engine.Context) error {
	if ctx == nil {
		return nil
	}
	conn := ctx.Connection
	if conn == nil {
		return nil
	}
	return conn.Close()
}
