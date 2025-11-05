package filters

import (
	"github.com/0xSplits/indexingo/client"
)

type Config struct {
	Cli client.Interface
}

type Filters struct {
	cli client.Interface
}

func New(c Config) *Filters {
	if c.Cli == nil {
		c.Cli = client.New(client.Config{})
	}

	return &Filters{
		cli: c.Cli,
	}
}
