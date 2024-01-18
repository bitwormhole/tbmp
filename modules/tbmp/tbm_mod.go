package tbmp

import (
	"github.com/bitwormhole/tbmp"
	"github.com/starter-go/application"
)

// ModuleForMain ...
func ModuleForMain() application.Module {

	mb := tbmp.NewMainModule()

	return mb.Create()
}
