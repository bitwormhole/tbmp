package tbmp

import (
	"github.com/bitwormhole/tbmp"
	"github.com/bitwormhole/tbmp/gen/client4tbmp"
	"github.com/bitwormhole/tbmp/gen/main4tbmp"
	"github.com/bitwormhole/tbmp/gen/server4tbmp"
	"github.com/bitwormhole/tbmp/gen/test4tbmp"
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

// ModuleForMain ...
func ModuleForMain() application.Module {

	mb := tbmp.NewMainModule()
	mb.Components(main4tbmp.ExportComponents)

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	m1 := ModuleForClient()
	m2 := ModuleForServer()

	mb := tbmp.NewTestModule()
	mb.Components(test4tbmp.ExportComponents)

	mb.Depend(m1, m2)

	return mb.Create()
}

// ModuleForClient ...
func ModuleForClient() application.Module {

	mb := tbmp.NewClientModule()
	mb.Components(client4tbmp.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

// ModuleForServer ...
func ModuleForServer() application.Module {

	mb := tbmp.NewServerModule()
	mb.Components(server4tbmp.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}
