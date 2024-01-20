package tbmp

// Option 定义选项函数的签名
type Option func(cfg *Configuration)

// OptionHeader 用于设置 header
func OptionHeader(name, value string) Option {
	return func(cfg *Configuration) {
		cfg.Headers.Set(name, value)
	}
}

// OptionMethod 用于设置 method
func OptionMethod(method string) Option {
	return func(cfg *Configuration) {
		cfg.Method = method
	}
}
