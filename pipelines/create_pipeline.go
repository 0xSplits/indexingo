package pipelines

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type CreatePipelineRequest struct {
	Name           string                        `json:"name"`
	Transformation string                        `json:"transformation"`
	Filter         string                        `json:"filter"`
	FilterKeys     []string                      `json:"filterKeys"`
	Networks       []string                      `json:"networks"`
	Enabled        bool                          `json:"enabled"`
	Delivery       CreatePipelineRequestDelivery `json:"delivery"`
}

type CreatePipelineRequestDelivery struct {
	Adapter    string                                  `json:"adapter"`
	Connection CreatePipelineRequestDeliveryConnection `json:"connection"`
}

type CreatePipelineRequestDeliveryConnection struct {
	Host    string            `json:"host"`
	Headers map[string]string `json:"headers,omitempty"`
}

type CreatePipelineResponse struct {
	// TODO no message is returned on success, but instead the pipeline structure itself !?
	Message string `json:"message"`
}

func (p *Pipelines) CreatePipeline(bod CreatePipelineRequest) (CreatePipelineResponse, error) {
	var err error

	var byt []byte
	{
		byt, err = json.Marshal(bod)
		if err != nil {
			return CreatePipelineResponse{}, tracer.Mask(err)
		}
	}

	var pat string
	{
		pat = p.cli.Endpoint() + "/dw/pipelines/"
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodPost, pat, bytes.NewReader(byt))
		if err != nil {
			return CreatePipelineResponse{}, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res CreatePipelineResponse
	{
		err = p.cli.Request(req, &res)
		if err != nil {
			return CreatePipelineResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
