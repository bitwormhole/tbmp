package main4tbmp
import (
    p027e886db "github.com/bitwormhole/tbmp/src/main/golang/code"
     "github.com/starter-go/application"
)

// type p027e886db.Example in package:github.com/bitwormhole/tbmp/src/main/golang/code
//
// id:com-027e886db07ff6ed-code-Example
// class:
// alias:
// scope:singleton
//
type p027e886db0_code_Example struct {
}

func (inst* p027e886db0_code_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-027e886db07ff6ed-code-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p027e886db0_code_Example) new() any {
    return &p027e886db.Example{}
}

func (inst* p027e886db0_code_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p027e886db.Example)
	nop(ie, com)

	


    return nil
}


