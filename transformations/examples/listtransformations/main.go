package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/transformations"
	"github.com/xh3b4sd/tracer"
)

//
//     go run transformations/examples/listtransformations/main.go
//

func main() {
	var err error

	var tra *transformations.Transformations
	{
		tra = transformations.New(transformations.Config{})
	}

	var res transformations.ListTransformationsResponse
	{
		res, err = tra.ListTransformations()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for _, x := range res.Data {
		fmt.Printf("%#v\n", x.Name)
	}
}
