package main

import (
	"os"

	"github.com/bitwormhole/tbmp/modules/tbmp"
	"github.com/starter-go/units"
)

func main() {

	m := tbmp.ModuleForTest()

	// i := starter.Init(os.Args)
	// i.MainModule(m)
	// i.WithPanic(true).Run()

	runner := units.NewRunner()
	runner.Dependencies(m)
	runner.EnablePanic(true).Run(os.Args)

}
