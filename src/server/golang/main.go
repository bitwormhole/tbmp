package main

import (
	"os"

	"github.com/bitwormhole/tbmp/modules/tbmp"
	"github.com/starter-go/starter"
)

func main() {
	m := tbmp.ModuleForMain()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
