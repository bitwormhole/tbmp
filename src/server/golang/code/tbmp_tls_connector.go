package code

import (
	"github.com/bitwormhole/tbmp/engine"
	"github.com/starter-go/application"
)

// TLSConnector ...
type TLSConnector struct {

	//starter:component
	_as func(engine.ElementRegistry) //starter:as(".")

	Host    string //starter:inject("${server.tbmps.host}")
	Port    int    //starter:inject("${server.tbmps.port}")
	Enabled bool   //starter:inject("${server.tbmps.enabled}")

	inner *serverSideConnector
}

func (inst *TLSConnector) _impl() (engine.ElementRegistry, application.Lifecycle) {
	return inst, inst
}

func (inst *TLSConnector) getInner() *serverSideConnector {
	i := inst.inner
	if i == nil {
		i = new(serverSideConnector)
		i.enabled = inst.Enabled
		i.config.Host = inst.Host
		i.config.Port = inst.Port
		inst.inner = i
	}
	return i
}

// Life ...
func (inst *TLSConnector) Life() *application.Life {
	return inst.getInner().Life()
}

// Elements ...
func (inst *TLSConnector) Elements() []*engine.ElementRegistration {
	return inst.getInner().Elements()
}
