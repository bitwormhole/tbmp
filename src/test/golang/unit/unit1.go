package unit

import (
	"github.com/bitwormhole/tbmp"
	"github.com/starter-go/units"
)

// Unit1 ... 单元测试示例
type Unit1 struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	Client tbmp.Client //starter:inject("#")
	Port   int         //starter:inject("${server.tbmp.port}")

}

func (inst *Unit1) _impl() units.Units { return inst }

// Units ...
func (inst *Unit1) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     "unit-1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	})

	return list
}

// Units ...
func (inst *Unit1) test1() error {

	cfg := inst.getConfig()
	h := inst.getHeaders()

	conn, err := inst.Client.Connect(cfg, h)
	if err != nil {
		return err
	}

	defer func() {
		conn.Close()
	}()

	tx, err := conn.Upstream()
	if err != nil {
		return err
	}

	rx, err := conn.Downstream()
	if err != nil {
		return err
	}

	rx.Headers()
	tx.Headers()
	return nil
}

func (inst *Unit1) getConfig() *tbmp.Configuration {
	cfg := &tbmp.Configuration{
		Port:     inst.Port,
		Host:     "localhost",
		Protocol: "tbmp",
	}
	return cfg
}

func (inst *Unit1) getHeaders() *tbmp.Headers {
	h := &tbmp.Headers{}
	h.Set("a", "b")
	return h
}
