package engine

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/tbmp"
)

// Context ...
type Context struct {
	Engine     *Engine
	Handler    Handler
	Selector   tbmp.HandlerSelector
	Connection tbmp.ServerSideConnection
	Upstream   tbmp.Upstream
	Downstream tbmp.Downstream
}

// Send 发送下行数据流的头部
func (inst *Context) Send(status int, contentType string) error {

	if status == 0 {
		status = http.StatusInternalServerError
	}
	statusText := strconv.Itoa(status) + " " + http.StatusText(status)

	down := &inst.Downstream
	down.ContentType = contentType
	down.Status = status

	headers := &inst.Downstream.Headers
	headers.Set(tbmp.HeaderContentType, down.ContentType)
	headers.Set(tbmp.HeaderStatus, statusText)

	conn := inst.Connection
	_, err := conn.Downstream(headers)
	return err
}
