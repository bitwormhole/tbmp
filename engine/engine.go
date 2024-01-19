package engine

// Engine ...
type Engine struct {
	Handlers []*HandlerRegistration
	Filters  []*FilterRegistration
	Elements []*ElementRegistration
	Chain    FilterChain
}
