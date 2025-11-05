package pipelines

import (
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type ListPipelinesResponse struct {
	Data []ListPipelinesResponseData `json:"data"`
}

type ListPipelinesResponseData struct {
	Name           string                        `json:"name"`
	Transformation string                        `json:"transformation"`
	Filter         string                        `json:"filter"`
	FilterKeys     []string                      `json:"filterKeys"`
	Networks       []string                      `json:"networks"`
	Enabled        bool                          `json:"enabled"`
	Delivery       ListPipelinesResponseDelivery `json:"delivery"`
}

type ListPipelinesResponseDelivery struct {
	Adapter    string                                  `json:"adapter"`
	Connection ListPipelinesResponseDeliveryConnection `json:"connection"`
}

type ListPipelinesResponseDeliveryConnection struct {
	Host    string            `json:"host"`
	Headers map[string]string `json:"headers,omitempty"`
}

// https://docs.indexing.co/guide/pipelines/list
func (p *Pipelines) ListPipelines() (ListPipelinesResponse, error) {
	var pat string
	{
		pat = p.cli.Endpoint() + "/dw/pipelines/"
	}

	req, err := http.NewRequest(http.MethodGet, pat, nil)
	if err != nil {
		return ListPipelinesResponse{}, tracer.Mask(err)
	}

	var val ListPipelinesResponse
	err = p.cli.Request(req, &val)
	if err != nil {
		return ListPipelinesResponse{}, tracer.Mask(err)
	}

	return val, nil
}
