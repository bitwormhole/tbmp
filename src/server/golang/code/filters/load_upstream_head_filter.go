package filters

import (
	"github.com/bitwormhole/tbmp"
	"github.com/bitwormhole/tbmp/engine"
)

// LoadUpstreamHeadFilter ...
type LoadUpstreamHeadFilter struct {

	//starter:component
	_as func(engine.FilterRegistry) //starter:as(".")

}

func (inst *LoadUpstreamHeadFilter) _impl() (engine.FilterRegistry, engine.Filter) {
	return inst, inst
}

// Filters ...
func (inst *LoadUpstreamHeadFilter) Filters() []*engine.FilterRegistration {
	r1 := &engine.FilterRegistration{
		Enabled: true,
		Order:   engine.FilterOrderLoadUpstreamHead,
		Filter:  inst,
	}
	return []*engine.FilterRegistration{r1}
}

// DoFilter ...
func (inst *LoadUpstreamHeadFilter) DoFilter(c *engine.Context, next engine.FilterChain) error {

	conn := c.Connection
	up, err := conn.Upstream()
	if err != nil {
		return err
	}
	h := up.Headers()
	c.Service = h.Get(tbmp.HeaderService)

	return next.DoFilter(c)
}
