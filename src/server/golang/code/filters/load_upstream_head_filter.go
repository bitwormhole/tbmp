package filters

import (
	"net/url"

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
	method := h.Get(tbmp.HeaderMethod)
	ctype := h.Get(tbmp.HeaderContentType)
	protocol := h.Get(tbmp.HeaderProtocol)

	// location
	location1 := h.Get(tbmp.HeaderURL)       // string
	location2, err := url.Parse(location1)   // url.URL
	location3 := tbmp.NewLocation(location2) // location
	if err != nil {
		return err
	}

	// selector
	hsb := &tbmp.HandlerSelectorBuilder{
		Method: method,
		Path:   location2.Path,
	}
	sel := hsb.Create()
	c.Selector = sel

	c.Upstream.Method = method
	c.Upstream.URL = location1
	c.Upstream.Location = *location3
	c.Upstream.ContentType = ctype
	c.Upstream.Headers = *h
	c.Upstream.Protocol = protocol

	return next.DoFilter(c)
}
