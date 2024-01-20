package code

import "github.com/bitwormhole/tbmp"

// ClientFacade ...
type ClientFacade struct {

	//starter:component
	_as func(tbmp.Client) //starter:as("#")

	inner tbmp.Client
}

func (inst *ClientFacade) _impl() tbmp.Client {
	return inst
}

func (inst *ClientFacade) getInner() tbmp.Client {
	c := inst.inner
	if c == nil {
		c = tbmp.NewClient()
		inst.inner = c
	}
	return c
}

// Name ...
func (inst *ClientFacade) Name() string {
	return inst.getInner().Name()
}

// SetName ...
func (inst *ClientFacade) SetName(name string) {
	inst.getInner().SetName(name)
}

// Connect ...
func (inst *ClientFacade) Connect(url string, options ...tbmp.Option) (tbmp.ClientSideConnection, error) {
	return inst.getInner().Connect(url, options...)
}
