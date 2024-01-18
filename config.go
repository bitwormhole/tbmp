package tbmp

import "strconv"

// 定义协议的名称（和版本）
const (
	NameProtocol = "TBMP/1.0"
	NameClient   = "TBMP.client/1.0"
	NameServer   = "TBMP.server/1.0"
)

// Configuration ...
type Configuration struct {
	Protocol string
	Host     string
	Port     int
	Username string
	Password string
	Secure   bool
	// Headers  Headers
}

func (cfg *Configuration) getAddress() string {
	if cfg == nil {
		cfg = new(Configuration)
	}
	const portDefault = 5570
	host := cfg.Host
	port := cfg.Port
	if port < 1 {
		port = portDefault
	}
	return host + ":" + strconv.Itoa(port)
}
