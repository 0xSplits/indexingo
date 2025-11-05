package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/filters"
	"github.com/xh3b4sd/tracer"
)

//
//     go run filters/examples/listfilters/main.go
//

func main() {
	var err error

	var fil *filters.Filters
	{
		fil = filters.New(filters.Config{})
	}

	var res filters.ListFiltersResponse
	{
		res, err = fil.ListFilters()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for _, x := range res.Data {
		fmt.Printf("%#v\n", x.Name)
	}
}
