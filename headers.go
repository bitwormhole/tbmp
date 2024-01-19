package tbmp

import "strings"

// 定义一些常用的头部字段名称
const (
	// HeaderMethod      = "Method"
	// HeaderPath        = "Path"

	HeaderService     = "Service"
	HeaderClient      = "Client"
	HeaderServer      = "Server"
	HeaderProtocol    = "Protocol"
	HeaderLocation    = "Location"
	HeaderContentType = "Content-Type"
)

// Headers ...
type Headers struct {
	table map[string]string
}

func (inst *Headers) normalname(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

// Get ...
func (inst *Headers) Get(name string) string {
	name = inst.normalname(name)
	t := inst.table
	if t == nil {
		return ""
	}
	return t[name]
}

// Set ...
func (inst *Headers) Set(name, value string) {
	name = inst.normalname(name)
	t := inst.table
	if t == nil {
		t = make(map[string]string)
		inst.table = t
	}
	t[name] = value
}

// ForEach ...
func (inst *Headers) ForEach(fn func(name, value string)) {
	t := inst.table
	if t == nil || fn == nil {
		return
	}
	for k, v := range t {
		fn(k, v)
	}
}
