package transformations

import (
	"github.com/0xSplits/indexingo/client"
)

type Config struct {
	Cli client.Interface
}

type Transformations struct {
	cli client.Interface
}

func New(c Config) *Transformations {
	if c.Cli == nil {
		c.Cli = client.New(client.Config{})
	}

	return &Transformations{
		cli: c.Cli,
	}
}
