package code

import (
	"github.com/bitwormhole/tbmp/engine"
	"github.com/starter-go/application"
)

// EngineBootstrap 这个组件用于控制默认的 TBMP 引擎
type EngineBootstrap struct {

	//starter:component

	FilterRegs  []engine.FilterRegistry  //starter:inject(".")
	HandlerRegs []engine.HandlerRegistry //starter:inject(".")
	ElementRegs []engine.ElementRegistry //starter:inject(".")

	engine *engine.Engine
}

func (inst *EngineBootstrap) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *EngineBootstrap) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.onCreate,
		OnStart:  inst.onStart,
		OnStop:   inst.onStop,
	}
}

func (inst *EngineBootstrap) onCreate() error {
	eng, err := inst.createEngine()
	if err != nil {
		return err
	}
	inst.engine = eng
	return nil
}

func (inst *EngineBootstrap) onStart() error {
	return nil
}

func (inst *EngineBootstrap) onStop() error {
	return nil
}

func (inst *EngineBootstrap) initElements(eng *engine.Engine) error {
	list := eng.Elements
	for _, item := range list {
		if item == nil {
			continue
		}
		if item.OnInit == nil || !item.Enabled {
			continue
		}
		err := item.OnInit(eng)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *EngineBootstrap) createEngine() (*engine.Engine, error) {

	eng := new(engine.Engine)

	eng.Elements = inst.loadElements()
	eng.Filters = inst.loadFilters()
	eng.Handlers = inst.loadHandlers()

	fcb := &engine.FilterChainBuilder{}
	fcb.AddRegistrations(eng.Filters...)
	eng.Chain = fcb.Create()

	err := inst.initElements(eng)
	if err != nil {
		return nil, err
	}
	return eng, nil
}

func (inst *EngineBootstrap) loadElements() []*engine.ElementRegistration {
	dst := make([]*engine.ElementRegistration, 0)
	src := inst.ElementRegs
	for _, r1 := range src {
		r2 := r1.Elements()
		dst = append(dst, r2...)
	}
	return dst
}

func (inst *EngineBootstrap) loadFilters() []*engine.FilterRegistration {
	dst := make([]*engine.FilterRegistration, 0)
	src := inst.FilterRegs
	for _, r1 := range src {
		r2 := r1.Filters()
		dst = append(dst, r2...)
	}
	return dst
}

func (inst *EngineBootstrap) loadHandlers() []*engine.HandlerRegistration {
	dst := make([]*engine.HandlerRegistration, 0)
	src := inst.HandlerRegs
	for _, r1 := range src {
		r2 := r1.Handlers()
		dst = append(dst, r2...)
	}
	return dst
}
