package engine

import (
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
