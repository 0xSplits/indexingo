package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/filters"
	"github.com/xh3b4sd/tracer"
)

//
//     go run filters/examples/addvalues/main.go
//

func main() {
	var err error

	var fil *filters.Filters
	{
		fil = filters.New(filters.Config{})
	}

	var res filters.AddValuesResponse
	{
		res, err = fil.AddValues("test-filter", []string{"0xb7f5bf799fb265657c628ef4a13f90f83a3a616a"})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		fmt.Printf("%#v\n", res.Message)
	}
}
