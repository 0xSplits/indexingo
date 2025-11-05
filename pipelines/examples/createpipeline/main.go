package main

import (
	"fmt"
	"os"

	"github.com/0xSplits/indexingo/pipelines"
	"github.com/xh3b4sd/tracer"
)

//
//     go run pipelines/examples/createpipeline/main.go
//

func main() {
	var err error

	var tra *pipelines.Pipelines
	{
		tra = pipelines.New(pipelines.Config{})
	}

	var bod pipelines.CreatePipelineRequest
	{
		bod = pipelines.CreatePipelineRequest{
			Name:           "test-pipeline",
			Transformation: "test-transformation",
			Filter:         "test-filter",
			FilterKeys:     []string{"from", "to"},
			Networks:       []string{"ethereum", "base"},
			Enabled:        true,
			Delivery: pipelines.CreatePipelineRequestDelivery{
				Adapter: "WEBSOCKET",
				Connection: pipelines.CreatePipelineRequestDeliveryConnection{
					Host: "https://pulsar.testing.splits.org/indexing",
					Headers: map[string]string{
						"Authorization": "Bearer " + os.Getenv("PULSAR_WEBSOCKET_SECRET"),
					},
				},
			},
		}
	}

	var res pipelines.CreatePipelineResponse
	{
		res, err = tra.CreatePipeline(bod)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		fmt.Printf("%#v\n", res.Message)
	}
}
