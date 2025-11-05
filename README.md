# indexingo

Golang client for the Indexing Co Neighborhood API. See https://docs.indexing.co.

### Examples

- [Filter Examples](./filters/examples/)
- [Pipeline Examples](./pipelines/examples/)
- [Transformation Examples](./transformations/examples/)

```golang
package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/filters"
	"github.com/xh3b4sd/tracer"
)

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
```
