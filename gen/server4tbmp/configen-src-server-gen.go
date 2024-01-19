package server4tbmp
import (
    p38ad25bca "github.com/bitwormhole/tbmp/engine"
    p2026f032f "github.com/bitwormhole/tbmp/src/server/golang/code"
     "github.com/starter-go/application"
)

// type p2026f032f.EngineControl in package:github.com/bitwormhole/tbmp/src/server/golang/code
//
// id:com-2026f032f9324ab6-code-EngineControl
// class:
// alias:
// scope:singleton
//
type p2026f032f9_code_EngineControl struct {
}

func (inst* p2026f032f9_code_EngineControl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2026f032f9324ab6-code-EngineControl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2026f032f9_code_EngineControl) new() any {
    return &p2026f032f.EngineControl{}
}

func (inst* p2026f032f9_code_EngineControl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2026f032f.EngineControl)
	nop(ie, com)

	
    com.FilterRegs = inst.getFilterRegs(ie)
    com.HandlerRegs = inst.getHandlerRegs(ie)
    com.ElementRegs = inst.getElementRegs(ie)


    return nil
}


func (inst*p2026f032f9_code_EngineControl) getFilterRegs(ie application.InjectionExt)[]p38ad25bca.FilterRegistry{
    dst := make([]p38ad25bca.FilterRegistry, 0)
    src := ie.ListComponents(".class-38ad25bcacda9374230da8811af7ccbc-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p38ad25bca.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p2026f032f9_code_EngineControl) getHandlerRegs(ie application.InjectionExt)[]p38ad25bca.HandlerRegistry{
    dst := make([]p38ad25bca.HandlerRegistry, 0)
    src := ie.ListComponents(".class-38ad25bcacda9374230da8811af7ccbc-HandlerRegistry")
    for _, item1 := range src {
        item2 := item1.(p38ad25bca.HandlerRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p2026f032f9_code_EngineControl) getElementRegs(ie application.InjectionExt)[]p38ad25bca.ElementRegistry{
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



// type p2026f032f.TCPConnector in package:github.com/bitwormhole/tbmp/src/server/golang/code
//
// id:com-2026f032f9324ab6-code-TCPConnector
// class:class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry
// alias:
// scope:singleton
//
type p2026f032f9_code_TCPConnector struct {
}

func (inst* p2026f032f9_code_TCPConnector) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2026f032f9324ab6-code-TCPConnector"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2026f032f9_code_TCPConnector) new() any {
    return &p2026f032f.TCPConnector{}
}

func (inst* p2026f032f9_code_TCPConnector) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2026f032f.TCPConnector)
	nop(ie, com)

	
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p2026f032f9_code_TCPConnector) getHost(ie application.InjectionExt)string{
    return ie.GetString("${server.tbmp.host}")
}


func (inst*p2026f032f9_code_TCPConnector) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.tbmp.port}")
}


func (inst*p2026f032f9_code_TCPConnector) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${server.tbmp.enabled}")
}



// type p2026f032f.TLSConnector in package:github.com/bitwormhole/tbmp/src/server/golang/code
//
// id:com-2026f032f9324ab6-code-TLSConnector
// class:class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry
// alias:
// scope:singleton
//
type p2026f032f9_code_TLSConnector struct {
}

func (inst* p2026f032f9_code_TLSConnector) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2026f032f9324ab6-code-TLSConnector"
	r.Classes = "class-38ad25bcacda9374230da8811af7ccbc-ElementRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2026f032f9_code_TLSConnector) new() any {
    return &p2026f032f.TLSConnector{}
}

func (inst* p2026f032f9_code_TLSConnector) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2026f032f.TLSConnector)
	nop(ie, com)

	
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p2026f032f9_code_TLSConnector) getHost(ie application.InjectionExt)string{
    return ie.GetString("${server.tbmps.host}")
}


func (inst*p2026f032f9_code_TLSConnector) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.tbmps.port}")
}


func (inst*p2026f032f9_code_TLSConnector) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${server.tbmps.enabled}")
}


