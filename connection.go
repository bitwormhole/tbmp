package tbmp

import "io"

// Connection 表示一个 TBMP 连接
type Connection interface {
	io.Closer
}

// ClientSideConnection ...
type ClientSideConnection interface {
	Connection

	Downstream() (RxStream, error)

	Upstream(input ...*Headers) (TxStream, error)
}

// ServerSideConnection ...
type ServerSideConnection interface {
	Connection

	Upstream() (RxStream, error)

	Downstream(input ...*Headers) (TxStream, error)
}
