package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/pipelines"
	"github.com/xh3b4sd/tracer"
)

//
//     curl --request POST \
//          --url https://app.indexing.co/dw/pipelines/test-pipeline/backfill \
//          --header 'Content-Type: application/json' \
//          --header "x-api-key: $INDEXINGCO_API_KEY" \
//          --data '{
//              "network": "BASE",
//              "value": "0xb7f5bf799fb265657c628ef4a13f90f83a3a616a",
//              "beatStart": 37740907
//          }'
//

//
//     go run pipelines/examples/backfillpipeline/main.go
//

func main() {
	var err error

	var pip *pipelines.Pipelines
	{
		pip = pipelines.New(pipelines.Config{})
	}

	var bod pipelines.BackfillPipelineRequest
	{
		bod = pipelines.BackfillPipelineRequest{
			Network:   "base",
			Value:     "0xb7f5bf799fb265657c628ef4a13f90f83a3a616a",
			BeatStart: 37740907,
		}
	}

	var res pipelines.BackfillPipelineResponse
	{
		res, err = pip.BackfillPipeline("test-pipeline", bod)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		fmt.Printf("%#v\n", res.Message)
	}
}
