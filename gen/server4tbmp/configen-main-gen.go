package server4tbmp

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p2026f032f9_code_EngineBootstrap{})
    inst.register(&p2026f032f9_code_Example{})
    inst.register(&p4c37db26f7_elements_TCPConnector{})
    inst.register(&p4c37db26f7_elements_TLSConnector{})
    inst.register(&p6e65738189_filters_FindHandlerFilter{})
    inst.register(&p6e65738189_filters_InvokeHandlerFilter{})
    inst.register(&p6e65738189_filters_LoadUpstreamHeadFilter{})
    inst.register(&pad76024e6d_handlers_Example{})


    return nil
}
