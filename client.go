package tbmp

// Connector 表示一个 TBMP 连接器
type Connector interface {
	Connect(url string, options ...Option) (ClientSideConnection, error)
}

// Client ...
type Client interface {
	Connector

	Name() string

	SetName(name string)
}

// NewClient ...
func NewClient() Client {
	return new(clientImpl)
}
