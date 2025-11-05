package pipelines

import (
	"github.com/0xSplits/indexingo/client"
)

type Config struct {
	Cli client.Interface
}

type Pipelines struct {
	cli client.Interface
}

func New(c Config) *Pipelines {
	if c.Cli == nil {
		c.Cli = client.New(client.Config{})
	}

	return &Pipelines{
		cli: c.Cli,
	}
}
