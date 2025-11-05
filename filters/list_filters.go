package filters

import (
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type ListFiltersResponse struct {
	Data []ListFiltersResponseFilter `json:"data"`
}

type ListFiltersResponseFilter struct {
	Name string `json:"name"`
}

// https://docs.indexing.co/guide/filters/list
func (f *Filters) ListFilters() (ListFiltersResponse, error) {
	var pat string
	{
		pat = f.cli.Endpoint() + "/dw/filters/"
	}

	req, err := http.NewRequest(http.MethodGet, pat, nil)
	if err != nil {
		return ListFiltersResponse{}, tracer.Mask(err)
	}

	var val ListFiltersResponse
	err = f.cli.Request(req, &val)
	if err != nil {
		return ListFiltersResponse{}, tracer.Mask(err)
	}

	return val, nil
}
