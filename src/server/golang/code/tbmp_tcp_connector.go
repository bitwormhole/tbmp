package code

import (
	"github.com/bitwormhole/tbmp/engine"
	"github.com/starter-go/application"
)

// TCPConnector ...
type TCPConnector struct {

	//starter:component
	_as func(engine.ElementRegistry) //starter:as(".")

	Host    string //starter:inject("${server.tbmp.host}")
	Port    int    //starter:inject("${server.tbmp.port}")
	Enabled bool   //starter:inject("${server.tbmp.enabled}")

	inner *serverSideConnector
}

func (inst *TCPConnector) _impl() (engine.ElementRegistry, application.Lifecycle) {
	return inst, inst
}

func (inst *TCPConnector) getInner() *serverSideConnector {
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
func (inst *TCPConnector) Life() *application.Life {
	return inst.getInner().Life()
}

// Elements ...
func (inst *TCPConnector) Elements() []*engine.ElementRegistration {
	return inst.getInner().Elements()
}
