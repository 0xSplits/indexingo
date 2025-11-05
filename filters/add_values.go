package filters

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type AddValuesResponse struct {
	Message string `json:"message"`
}

type addValuesRequest struct {
	Values []string `json:"values"`
}

// https://docs.indexing.co/guide/filters/add
func (f *Filters) AddValues(nam string, val []string) (AddValuesResponse, error) {
	var err error

	var bod addValuesRequest
	{
		bod = addValuesRequest{
			Values: val,
		}
	}

	var pat string
	{
		pat = f.cli.Endpoint() + "/dw/filters/" + nam
	}

	var byt []byte
	{
		byt, err = json.Marshal(bod)
		if err != nil {
			return AddValuesResponse{}, tracer.Mask(err)
		}
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodPost, pat, bytes.NewReader(byt))
		if err != nil {
			return AddValuesResponse{}, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res AddValuesResponse
	{
		err = f.cli.Request(req, &res)
		if err != nil {
			return AddValuesResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
