package filters

import (
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type ListValuesResponse struct {
	Data     []string                   `json:"data"`
	Metadata ListValuesResponseMetadata `json:"metadata"`
}

type ListValuesResponseMetadata struct {
	NextPageToken *string `json:"nextPageToken"`
}

// https://docs.indexing.co/guide/filters/list
func (f *Filters) ListValues(nam string) (ListValuesResponse, error) {
	var pat string
	{
		pat = f.cli.Endpoint() + "/dw/filters/" + nam
	}

	req, err := http.NewRequest(http.MethodGet, pat, nil)
	if err != nil {
		return ListValuesResponse{}, tracer.Mask(err)
	}

	var val ListValuesResponse
	err = f.cli.Request(req, &val)
	if err != nil {
		return ListValuesResponse{}, tracer.Mask(err)
	}

	return val, nil
}
