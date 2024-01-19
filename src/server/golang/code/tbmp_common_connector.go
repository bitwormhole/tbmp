package code

import (
	"fmt"
	"io"
	"time"

	"github.com/bitwormhole/tbmp"
	"github.com/bitwormhole/tbmp/engine"
	"github.com/starter-go/application"
)

type serverSideConnector struct {
	config tbmp.Configuration
	engine *engine.Engine

	enabled  bool
	started  bool
	starting bool
	stopping bool
	stopped  bool

	closer io.Closer
}

// Life ...
func (inst *serverSideConnector) Life() *application.Life {

	if !inst.enabled {
		return &application.Life{}
	}

	return &application.Life{
		OnStart:    inst.onStart,
		OnStop:     inst.onStop,
		OnStopPost: inst.onStopPost,
		OnStopPre:  inst.onStopPre,
	}
}

// Elements ...
func (inst *serverSideConnector) Elements() []*engine.ElementRegistration {

	if !inst.enabled {
		return []*engine.ElementRegistration{}
	}

	r1 := &engine.ElementRegistration{
		Enabled: true,
		OnInit:  inst.onInit,
	}

	return []*engine.ElementRegistration{r1}
}

func (inst *serverSideConnector) onInit(e *engine.Engine) error {
	inst.engine = e
	return nil
}

func (inst *serverSideConnector) onStart() error {

	inst.starting = true

	runner := &serverSideConnRunner{parent: inst}

	go func() {
		runner.run1()
	}()

	return nil
}

func (inst *serverSideConnector) onStop() error {

	c := inst.closer
	inst.closer = nil
	inst.stopping = true

	if c != nil {
		go func() {
			err := c.Close()
			engine.HandleError(err)
		}()
	}
	return nil
}

func (inst *serverSideConnector) onStopPre() error {
	return inst.join()
}

func (inst *serverSideConnector) onStopPost() error {
	return inst.join()
}

func (inst *serverSideConnector) join() error {
	if !inst.starting {
		return fmt.Errorf("tbmp::serverSideConnector: no starting")
	}
	for {
		if inst.stopped {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type serverSideConnRunner struct {
	parent *serverSideConnector
}

func (inst *serverSideConnRunner) config() *tbmp.Configuration {
	return &inst.parent.config
}

func (inst *serverSideConnRunner) run1() {

	defer func() {
		x := recover()
		engine.HandleErrorX(x)
	}()

	defer func() {
		inst.parent.stopped = true
	}()

	err := inst.run2()
	if err != nil {
		engine.HandleError(err)
	}
}

func (inst *serverSideConnRunner) run2() error {

	cfg := inst.config()
	ser := tbmp.NewServer()

	l, err := ser.Listen(cfg)
	if err != nil {
		return err
	}

	inst.parent.started = true

	for {
		if inst.parent.stopping {
			break
		}
		conn, err := l.Accept()
		if err == nil {
			inst.onConnect1(conn)
		} else {
			engine.HandleError(err)
		}
	}

	return nil
}

func (inst *serverSideConnRunner) onConnect1(conn tbmp.ServerSideConnection) {

	defer func() {
		x := recover()
		engine.HandleErrorX(x)
	}()

	ctx := new(engine.Context)
	ctx.Connection = conn
	ctx.Engine = inst.parent.engine
	ctx.Handler = nil

	go func() {
		inst.onConnect2(ctx)
	}()
}

func (inst *serverSideConnRunner) onConnect2(ctx *engine.Context) {
	defer func() {
		x := recover()
		engine.HandleErrorX(x)
	}()
	err := ctx.Engine.Chain.DoFilter(ctx)
	engine.HandleError(err)
}
