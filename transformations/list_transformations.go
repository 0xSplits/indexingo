package transformations

import (
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type ListTransformationsResponse struct {
	Data []ListTransformationsResponseFilter `json:"data"`
}

type ListTransformationsResponseFilter struct {
	Name string `json:"name"`
}

// https://docs.indexing.co/guide/transformations/list
func (t *Transformations) ListTransformations() (ListTransformationsResponse, error) {
	var pat string
	{
		pat = t.cli.Endpoint() + "/dw/transformations/"
	}

	req, err := http.NewRequest(http.MethodGet, pat, nil)
	if err != nil {
		return ListTransformationsResponse{}, tracer.Mask(err)
	}

	var val ListTransformationsResponse
	err = t.cli.Request(req, &val)
	if err != nil {
		return ListTransformationsResponse{}, tracer.Mask(err)
	}

	return val, nil
}
