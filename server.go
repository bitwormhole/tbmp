package tbmp

import (
	"io"
)

// Listener ...
type Listener interface {
	io.Closer

	Accept() (ServerSideConnection, error)
}

// Server ...
type Server interface {
	Name() string

	SetName(name string)

	Listen(cfg *Configuration) (Listener, error)
}

// NewServer ...
func NewServer() Server {
	server := new(serverImpl)
	return server
}
