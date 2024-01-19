package engine

import "github.com/bitwormhole/tbmp"

// Context ...
type Context struct {
	Engine     *Engine
	Connection tbmp.ServerSideConnection
	Handler    Handler
	Service    string // service name of connection
}
