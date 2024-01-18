package tbmp

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
)

type connectionContext struct {
	config    *Configuration
	conn      net.Conn
	closer    io.Closer
	reader    io.Reader
	writer    io.Writer
	rxHdr     Headers
	txHdr     Headers
	rxHdrDone bool
	txHdrDone bool
}

func (inst *connectionContext) sendHead() error {

	if inst.txHdrDone {
		return nil
	}
	inst.txHdrDone = true

	b := bytes.NewBuffer(nil)
	h := &inst.txHdr
	h.ForEach(func(name, value string) {
		b.WriteString(name)
		b.WriteString(":")
		b.WriteString(value)
		b.WriteString("\n")
	})
	b.WriteByte(0)
	data := b.Bytes()
	cb, err := inst.writer.Write(data)
	if err != nil {
		return err
	}
	if cb != len(data) {
		return fmt.Errorf("bad head size")
	}
	return nil
}

func (inst *connectionContext) receiveHead() error {

	if inst.rxHdrDone {
		return nil
	}
	inst.rxHdrDone = true

	buffer1 := make([]byte, 1)
	buffer2 := bytes.NewBuffer(nil)
	h := &inst.rxHdr
	conn := inst.reader
	for {
		cb, err := conn.Read(buffer1)
		if err != nil {
			return err
		}
		if cb != 1 {
			return fmt.Errorf("bad size")
		}
		b := buffer1[0]
		if b == 0 {
			break
		} else if b == '\n' {
			inst.parseHeader(buffer2, h)
			buffer2.Reset()
		} else {
			buffer2.WriteByte(b)
		}
	}
	// inst.parseHeader(buffer2, h)
	return nil
}

func (inst *connectionContext) parseHeader(src *bytes.Buffer, dst *Headers) {
	bin := src.Bytes()
	str := string(bin)
	i := strings.IndexByte(str, ':')
	name := strings.TrimSpace(str[0:i])
	value := strings.TrimSpace(str[i+1:])
	dst.Set(name, value)
}

func (inst *connectionContext) Close() error {
	c := inst.closer
	inst.closer = nil
	if c == nil {
		return nil
	}
	return c.Close()
}
