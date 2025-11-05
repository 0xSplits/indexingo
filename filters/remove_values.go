package filters

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type RemoveValuesResponse struct {
	Message string `json:"message"`
}

type removeValuesRequest struct {
	Values []string `json:"values"`
}

// https://docs.indexing.co/guide/filters/remove
func (f *Filters) RemoveValues(nam string, val []string) (RemoveValuesResponse, error) {
	var err error

	var bod removeValuesRequest
	{
		bod = removeValuesRequest{
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
			return RemoveValuesResponse{}, tracer.Mask(err)
		}
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodDelete, pat, bytes.NewReader(byt))
		if err != nil {
			return RemoveValuesResponse{}, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res RemoveValuesResponse
	{
		err = f.cli.Request(req, &res)
		if err != nil {
			return RemoveValuesResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
