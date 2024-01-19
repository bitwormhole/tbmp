package test4tbmp
import (
    pb198020f1 "github.com/bitwormhole/tbmp"
    p8505ed25c "github.com/bitwormhole/tbmp/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p8505ed25c.DemoUnit in package:github.com/bitwormhole/tbmp/src/test/golang/unit
//
// id:com-8505ed25cac7a1c5-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p8505ed25ca_unit_DemoUnit struct {
}

func (inst* p8505ed25ca_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8505ed25cac7a1c5-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8505ed25ca_unit_DemoUnit) new() any {
    return &p8505ed25c.DemoUnit{}
}

func (inst* p8505ed25ca_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8505ed25c.DemoUnit)
	nop(ie, com)

	


    return nil
}



// type p8505ed25c.Unit1 in package:github.com/bitwormhole/tbmp/src/test/golang/unit
//
// id:com-8505ed25cac7a1c5-unit-Unit1
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p8505ed25ca_unit_Unit1 struct {
}

func (inst* p8505ed25ca_unit_Unit1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8505ed25cac7a1c5-unit-Unit1"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8505ed25ca_unit_Unit1) new() any {
    return &p8505ed25c.Unit1{}
}

func (inst* p8505ed25ca_unit_Unit1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8505ed25c.Unit1)
	nop(ie, com)

	
    com.Client = inst.getClient(ie)
    com.Port = inst.getPort(ie)


    return nil
}


func (inst*p8505ed25ca_unit_Unit1) getClient(ie application.InjectionExt)pb198020f1.Client{
    return ie.GetComponent("#alias-b198020f1a6c075088f46456e12cfd75-Client").(pb198020f1.Client)
}


func (inst*p8505ed25ca_unit_Unit1) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.tbmp.port}")
}


