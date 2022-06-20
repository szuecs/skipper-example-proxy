package filters

import (
	"github.com/zalando/skipper/filters"
)

type myFilter struct{}

func NewMyFilter() *myFilter {
	return &myFilter{}
}

func (spec *myFilter) Name() string { return "myFilter" }

func (spec *myFilter) CreateFilter(config []interface{}) (filters.Filter, error) {
	return NewMyFilter(), nil
}

func (f *myFilter) Request(ctx filters.FilterContext) {
	ctx.Request().Header.Set("my-filter", "request")
}

func (f *myFilter) Response(ctx filters.FilterContext) {
	ctx.Response().Header.Set("my-filter", "response")
}
