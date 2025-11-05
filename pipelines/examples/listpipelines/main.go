package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/pipelines"
	"github.com/xh3b4sd/tracer"
)

//
//     go run pipelines/examples/listpipelines/main.go
//

func main() {
	var err error

	var pip *pipelines.Pipelines
	{
		pip = pipelines.New(pipelines.Config{})
	}

	var res pipelines.ListPipelinesResponse
	{
		res, err = pip.ListPipelines()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for _, x := range res.Data {
		fmt.Printf("%#v\n", x.Name)
		fmt.Printf("%#v\n", x.Transformation)
		fmt.Printf("%#v\n", x.Filter)
		fmt.Printf("%#v\n", x.Delivery.Adapter)
		fmt.Printf("%#v\n", x.Delivery.Connection.Host)
	}
}
