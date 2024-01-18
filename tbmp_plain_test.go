package tbmp

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/starter-go/vlog"
)

func TestPlain(t *testing.T) {
	pt := &myPlainTester{}
	err := pt.run()
	if err != nil {
		t.Error(err)
	}
}

type myPlainTester struct {
	listener Listener

	starting bool
	stopping bool
	started  bool
	stopped  bool
}

func (inst *myPlainTester) config() *Configuration {
	c := &Configuration{
		Host: "",
		Port: 0,
	}
	return c
}

func (inst *myPlainTester) run() error {
	interval := time.Second * 3
	steps := make([]func() error, 0)
	steps = append(steps, inst.startServer)
	steps = append(steps, inst.openClient)
	steps = append(steps, inst.stopServer)
	steps = append(steps, inst.join)
	for _, fn := range steps {
		err := fn()
		if err != nil {
			return err
		}
		time.Sleep(interval)
	}
	return nil
}

func (inst *myPlainTester) runServer() error {

	defer func() {
		inst.stopped = true
	}()

	inst.log("run server")

	cfg := inst.config()
	cfg.Host = "0.0.0.0"

	server := NewServer()
	l, err := server.Listen(cfg)
	if err != nil {
		return err
	}

	defer func() {
		if l != nil {
			l.Close()
		}
	}()

	inst.started = true
	conn, err := l.Accept()
	if err != nil {
		return err
	}

	h1, err := conn.Upstream()
	if err != nil {
		return err
	}

	hx := &Headers{}
	hx.Set("foo", "bar(server-side)")
	h2, err := conn.Downstream(hx)
	if err != nil {
		return err
	}

	inst.log("server.up:   ", h1.Headers())
	inst.log("server.down: ", h2.Headers())

	inst.listener = l
	l = nil
	return nil
}

func (inst *myPlainTester) startServer() error {

	inst.log("start server")
	inst.starting = true

	go func() {
		err := inst.runServer()
		if err != nil {
			vlog.Error(err.Error())
		}
	}()

	return nil
}

func (inst *myPlainTester) openClient() error {
	inst.log("open client")

	cfg := inst.config()
	cfg.Host = "localhost"

	client := NewClient()
	conn, err := client.Connect(cfg)
	if err != nil {
		return err
	}

	defer func() {
		conn.Close()
	}()

	hx := &Headers{}
	hx.Set("foo", "bar(client-side)")
	h1, err := conn.Upstream(hx)
	if err != nil {
		return err
	}

	h2, err := conn.Downstream()
	if err != nil {
		return err
	}

	inst.log(h1.Headers())
	inst.log(h2.Headers())

	return nil
}

func (inst *myPlainTester) stopServer() error {
	inst.log("stop server")

	inst.stopping = true
	l := inst.listener
	inst.listener = nil
	if l == nil {
		return nil
	}
	return l.Close()
}

func (inst *myPlainTester) join() error {
	if !inst.starting {
		return fmt.Errorf("no start")
	}
	for {
		if inst.stopped {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (inst *myPlainTester) log(v ...any) {
	fb := strings.Builder{}
	for i := range v {
		if i > 0 {
			fb.WriteRune(' ')
		}
		fb.WriteString("%s")
	}
	f := fb.String()
	vlog.Info(f, v...)
}
