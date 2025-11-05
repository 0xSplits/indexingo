package pipelines

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xh3b4sd/tracer"
)

type BackfillPipelineRequest struct {
	BeatEnd   int64  `json:"beatEnd,omitempty"`
	BeatStart int64  `json:"beatStart,omitempty"`
	Network   string `json:"network,omitempty"`
	Value     string `json:"value,omitempty"`
}

type BackfillPipelineResponse struct {
	Message string `json:"message"`
}

// BackfillPipeline triggers a backfill for an existing pipeline by name. Note
// that this does only work for a single address per call at the time of
// writing. See also https://docs.indexing.co/guide/pipelines/backfill.
func (p *Pipelines) BackfillPipeline(nam string, bod BackfillPipelineRequest) (BackfillPipelineResponse, error) {
	var err error

	var byt []byte
	{
		byt, err = json.Marshal(bod)
		if err != nil {
			return BackfillPipelineResponse{}, tracer.Mask(err)
		}
	}

	var pat string
	{
		pat = p.cli.Endpoint() + "/dw/pipelines/" + nam + "/backfill"
	}

	var req *http.Request
	{
		req, err = http.NewRequest(http.MethodPost, pat, bytes.NewReader(byt))
		if err != nil {
			return BackfillPipelineResponse{}, tracer.Mask(err)
		}
	}

	{
		req.Header.Set("content-type", "application/json")
	}

	var res BackfillPipelineResponse
	{
		err = p.cli.Request(req, &res)
		if err != nil {
			return BackfillPipelineResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
