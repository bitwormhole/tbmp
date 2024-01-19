package engine

import "github.com/starter-go/base/util"

// 定义各个过滤器的处理顺序
const (
	FilterOrder0 = iota

	FilterOrderLoadUpstreamHead
	FilterOrderFindHandler
	FilterOrderInvokeHandler

	FilterOrder99
)

////////////////////////////////////////////////////////////////////////////////

// Filter ...
type Filter interface {
	DoFilter(c *Context, next FilterChain) error
}

// FilterChain ...
type FilterChain interface {
	DoFilter(c *Context) error
}

// FilterRegistration ...
type FilterRegistration struct {
	Enabled bool
	Order   int
	Filter  Filter
}

// FilterRegistry ...
type FilterRegistry interface {
	Filters() []*FilterRegistration
}

////////////////////////////////////////////////////////////////////////////////

// FilterChainBuilder ...
type FilterChainBuilder struct {
	items []*FilterRegistration
}

// Create ...
func (inst *FilterChainBuilder) Create() FilterChain {

	items := inst.prepareItems()
	end := new(filterChainEnding)
	var chain FilterChain = end

	for _, item := range items {
		node := new(filterChainNode)
		node.next = chain
		node.filter = item.Filter
		chain = node
	}

	return chain
}

// AddRegistrations ...
func (inst *FilterChainBuilder) AddRegistrations(src ...*FilterRegistration) *FilterChainBuilder {
	if src != nil {
		inst.items = append(inst.items, src...)
	}
	return inst
}

// AddRegistries ...
func (inst *FilterChainBuilder) AddRegistries(src ...FilterRegistry) *FilterChainBuilder {
	for _, r1 := range src {
		r2 := r1.Filters()
		inst.AddRegistrations(r2...)
	}
	return inst
}

func (inst *FilterChainBuilder) prepareItems() []*FilterRegistration {

	src := inst.items
	dst := make([]*FilterRegistration, 0)
	for _, reg := range src {
		if reg == nil {
			continue
		} else if !reg.Enabled || reg.Filter == nil {
			continue
		}
		dst = append(dst, reg)
	}

	// sort
	s := &util.Sorter{
		OnLen:  func() int { return len(dst) },
		OnSwap: func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] },
		OnLess: func(i1, i2 int) bool {
			o1 := dst[i1].Order
			o2 := dst[i2].Order
			return o1 < o2
		},
	}
	s.Sort()

	return dst
}

////////////////////////////////////////////////////////////////////////////////

type filterChainEnding struct{}

func (inst *filterChainEnding) _impl() FilterChain {
	return inst
}

func (inst *filterChainEnding) DoFilter(c *Context) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type filterChainNode struct {
	filter Filter
	next   FilterChain
}

func (inst *filterChainNode) _impl() FilterChain {
	return inst
}

func (inst *filterChainNode) DoFilter(c *Context) error {
	return inst.filter.DoFilter(c, inst.next)
}

////////////////////////////////////////////////////////////////////////////////
