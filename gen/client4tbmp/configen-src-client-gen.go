package client4tbmp
import (
    pc57e5c248 "github.com/bitwormhole/tbmp/src/client/golang/code"
     "github.com/starter-go/application"
)

// type pc57e5c248.ClientFacade in package:github.com/bitwormhole/tbmp/src/client/golang/code
//
// id:com-c57e5c248ce10a07-code-ClientFacade
// class:
// alias:alias-b198020f1a6c075088f46456e12cfd75-Client
// scope:singleton
//
type pc57e5c248c_code_ClientFacade struct {
}

func (inst* pc57e5c248c_code_ClientFacade) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c57e5c248ce10a07-code-ClientFacade"
	r.Classes = ""
	r.Aliases = "alias-b198020f1a6c075088f46456e12cfd75-Client"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc57e5c248c_code_ClientFacade) new() any {
    return &pc57e5c248.ClientFacade{}
}

func (inst* pc57e5c248c_code_ClientFacade) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc57e5c248.ClientFacade)
	nop(ie, com)

	


    return nil
}



// type pc57e5c248.Example in package:github.com/bitwormhole/tbmp/src/client/golang/code
//
// id:com-c57e5c248ce10a07-code-Example
// class:
// alias:
// scope:singleton
//
type pc57e5c248c_code_Example struct {
}

func (inst* pc57e5c248c_code_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c57e5c248ce10a07-code-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc57e5c248c_code_Example) new() any {
    return &pc57e5c248.Example{}
}

func (inst* pc57e5c248c_code_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc57e5c248.Example)
	nop(ie, com)

	


    return nil
}


