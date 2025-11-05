package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/filters"
	"github.com/xh3b4sd/tracer"
)

//
//     go run filters/examples/listvalues/main.go
//

func main() {
	var err error

	var fil *filters.Filters
	{
		fil = filters.New(filters.Config{})
	}

	var res filters.ListValuesResponse
	{
		res, err = fil.ListValues("test-filter")
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for _, x := range res.Data {
		fmt.Printf("%#v\n", x)
	}
}
