package main

import (
	"fmt"

	"github.com/0xSplits/indexingo/transformations"
	"github.com/xh3b4sd/tracer"
)

//
//     go run transformations/examples/createtransformation/main.go
//

func main() {
	var err error

	var tra *transformations.Transformations
	{
		tra = transformations.New(transformations.Config{})
	}

	var cod string
	{
		cod = `
		  function txfersByBlock(blo) {
			  const tra = templates.tokenTransfers(blo);

				return tra.map(txfer => ({
				  network: blo._network,
					chainId: utils.evmChainToId(blo._network),
					blockHash: blo.hash,
					blockNumber: blo.number,
					timestamp: utils.blockToTimestamp(blo),
					...tra,
				}));
			}
		`
	}

	var res transformations.CreateTransformationResponse
	{
		res, err = tra.CreateTransformation("test-transformation", cod)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		fmt.Printf("%#v\n", res.Message)
	}
}
