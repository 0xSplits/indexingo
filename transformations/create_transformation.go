package transformations

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type CreateTransformationResponse struct {
	Message string `json:"message"`
}

type createTransformationRequest struct {
	Code string `json:"code"`
}

// https://docs.indexing.co/guide/transformations/create
func (f *Transformations) CreateTransformation(nam string, cod string) (CreateTransformationResponse, error) {
	var err error

	var bod createTransformationRequest
	{
		bod = createTransformationRequest{
			Code: cod,
		}
	}

	var pat string
	{
		pat = f.cli.Endpoint() + "/dw/transformations/" + nam
	}

	var byt []byte
	{
		byt, err = json.Marshal(bod)
		if err != nil {
			return CreateTransformationResponse{}, tracer.Mask(err)
		}
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodPost, pat, bytes.NewReader(byt))
		if err != nil {
			return CreateTransformationResponse{}, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res CreateTransformationResponse
	{
		err = f.cli.Request(req, &res)
		if err != nil {
			return CreateTransformationResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
