package server4tbmp
import (
    p38ad25bca "github.com/bitwormhole/tbmp/engine"
    p2026f032f "github.com/bitwormhole/tbmp/src/server/golang/code"
    p4c37db26f "github.com/bitwormhole/tbmp/src/server/golang/code/elements"
    p6e6573818 "github.com/bitwormhole/tbmp/src/server/golang/code/filters"
    pad76024e6 "github.com/bitwormhole/tbmp/src/server/golang/code/handlers"
     "github.com/starter-go/application"
)

// type p2026f032f.EngineBootstrap in package:github.com/bitwormhole/tbmp/src/server/golang/code
//
// id:com-2026f032f9324ab6-code-EngineBootstrap
// class:
// alias:
// scope:singleton
//
type p2026f032f9_code_EngineBootstrap struct {
}

func (inst* p2026f032f9_code_EngineBootstrap) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2026f032f9324ab6-code-EngineBootstrap"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2026f032f9_code_EngineBootstrap) new() any {
    return &p2026f032f.EngineBootstrap{}
}

func (inst* p2026f032f9_code_EngineBootstrap) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2026f032f.EngineBootstrap)
	nop(ie, com)

	
    com.FilterRegs = inst.getFilterRegs(ie)
    com.HandlerRegs = inst.getHandlerRegs(ie)
    com.ElementRegs = inst.getElementRegs(ie)


    return nil
}


func (inst*p2026f032f9_code_EngineBootstrap) getFilterRegs(ie application.InjectionExt)[]p38ad25bca.FilterRegistry{
    dst := make([]p38ad25bca.FilterRegistry, 0)
    src := ie.ListComponents(".class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p38ad25bca.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p2026f032f9_code_EngineBootstrap) getHandlerRegs(ie application.InjectionExt)[]p38ad25bca.HandlerRegistry{
    dst := make([]p38ad25bca.HandlerRegistry, 0)
    src := ie.ListComponents(".class-38ad25bcacda9374230da8811af7ccbc-HandlerRegistry")
    for _, item1 := range src {
        item2 := item1.(p38ad25bca.HandlerRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p2026f032f9_code_EngineBootstrap) getElementRegs(ie application.InjectionExt)[]p38ad25bca.ElementRegistry{
    dst := make([]p38ad25bca.ElementRegistry, 0)
    src := ie.ListComponents(".class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry")
    for _, item1 := range src {
        item2 := item1.(p38ad25bca.ElementRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p2026f032f.Example in package:github.com/bitwormhole/tbmp/src/server/golang/code
//
// id:com-2026f032f9324ab6-code-Example
// class:
// alias:
// scope:singleton
//
type p2026f032f9_code_Example struct {
}

func (inst* p2026f032f9_code_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2026f032f9324ab6-code-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2026f032f9_code_Example) new() any {
    return &p2026f032f.Example{}
}

func (inst* p2026f032f9_code_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2026f032f.Example)
	nop(ie, com)

	


    return nil
}



// type p4c37db26f.TCPConnector in package:github.com/bitwormhole/tbmp/src/server/golang/code/elements
//
// id:com-4c37db26f7020ef9-elements-TCPConnector
// class:class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry
// alias:
// scope:singleton
//
type p4c37db26f7_elements_TCPConnector struct {
}

func (inst* p4c37db26f7_elements_TCPConnector) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4c37db26f7020ef9-elements-TCPConnector"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4c37db26f7_elements_TCPConnector) new() any {
    return &p4c37db26f.TCPConnector{}
}

func (inst* p4c37db26f7_elements_TCPConnector) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4c37db26f.TCPConnector)
	nop(ie, com)

	
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p4c37db26f7_elements_TCPConnector) getHost(ie application.InjectionExt)string{
    return ie.GetString("${server.tbmp.host}")
}


func (inst*p4c37db26f7_elements_TCPConnector) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.tbmp.port}")
}


func (inst*p4c37db26f7_elements_TCPConnector) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${server.tbmp.enabled}")
}



// type p4c37db26f.TLSConnector in package:github.com/bitwormhole/tbmp/src/server/golang/code/elements
//
// id:com-4c37db26f7020ef9-elements-TLSConnector
// class:class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry
// alias:
// scope:singleton
//
type p4c37db26f7_elements_TLSConnector struct {
}

func (inst* p4c37db26f7_elements_TLSConnector) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4c37db26f7020ef9-elements-TLSConnector"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4c37db26f7_elements_TLSConnector) new() any {
    return &p4c37db26f.TLSConnector{}
}

func (inst* p4c37db26f7_elements_TLSConnector) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4c37db26f.TLSConnector)
	nop(ie, com)

	
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p4c37db26f7_elements_TLSConnector) getHost(ie application.InjectionExt)string{
    return ie.GetString("${server.tbmps.host}")
}


func (inst*p4c37db26f7_elements_TLSConnector) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.tbmps.port}")
}


func (inst*p4c37db26f7_elements_TLSConnector) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${server.tbmps.enabled}")
}



// type p6e6573818.FindHandlerFilter in package:github.com/bitwormhole/tbmp/src/server/golang/code/filters
//
// id:com-6e657381895e3878-filters-FindHandlerFilter
// class:class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry
// alias:
// scope:singleton
//
type p6e65738189_filters_FindHandlerFilter struct {
}

func (inst* p6e65738189_filters_FindHandlerFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6e657381895e3878-filters-FindHandlerFilter"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6e65738189_filters_FindHandlerFilter) new() any {
    return &p6e6573818.FindHandlerFilter{}
}

func (inst* p6e65738189_filters_FindHandlerFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6e6573818.FindHandlerFilter)
	nop(ie, com)

	


    return nil
}



// type p6e6573818.InvokeHandlerFilter in package:github.com/bitwormhole/tbmp/src/server/golang/code/filters
//
// id:com-6e657381895e3878-filters-InvokeHandlerFilter
// class:class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry
// alias:
// scope:singleton
//
type p6e65738189_filters_InvokeHandlerFilter struct {
}

func (inst* p6e65738189_filters_InvokeHandlerFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6e657381895e3878-filters-InvokeHandlerFilter"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6e65738189_filters_InvokeHandlerFilter) new() any {
    return &p6e6573818.InvokeHandlerFilter{}
}

func (inst* p6e65738189_filters_InvokeHandlerFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6e6573818.InvokeHandlerFilter)
	nop(ie, com)

	


    return nil
}



// type p6e6573818.LoadUpstreamHeadFilter in package:github.com/bitwormhole/tbmp/src/server/golang/code/filters
//
// id:com-6e657381895e3878-filters-LoadUpstreamHeadFilter
// class:class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry
// alias:
// scope:singleton
//
type p6e65738189_filters_LoadUpstreamHeadFilter struct {
}

func (inst* p6e65738189_filters_LoadUpstreamHeadFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6e657381895e3878-filters-LoadUpstreamHeadFilter"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6e65738189_filters_LoadUpstreamHeadFilter) new() any {
    return &p6e6573818.LoadUpstreamHeadFilter{}
}

func (inst* p6e65738189_filters_LoadUpstreamHeadFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6e6573818.LoadUpstreamHeadFilter)
	nop(ie, com)

	


    return nil
}



// type pad76024e6.Example in package:github.com/bitwormhole/tbmp/src/server/golang/code/handlers
//
// id:com-ad76024e6dee68ec-handlers-Example
// class:
// alias:
// scope:singleton
//
type pad76024e6d_handlers_Example struct {
}

func (inst* pad76024e6d_handlers_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad76024e6dee68ec-handlers-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad76024e6d_handlers_Example) new() any {
    return &pad76024e6.Example{}
}

func (inst* pad76024e6d_handlers_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad76024e6.Example)
	nop(ie, com)

	


    return nil
}


