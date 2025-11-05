package client

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Cli is the optional HTTP client being used to handle requests and
	// responses. Defaults to http.DefaultClient.
	Cli *http.Client

	// Key is the required API key to make API interactions work. If this value is
	// empty, then the client will try to find an API key within the process
	// environment using the $INDEXINGCO_API_KEY environment variable. If that
	// lookup does then not yield a non-empty string, then the client creation
	// will panic.
	Key string

	// Url is the optional API endpoint. Defaults to "https://app.indexing.co".
	Url string
}

type Client struct {
	cli *http.Client
	key string
	url string
}

func New(c Config) *Client {
	if c.Cli == nil {
		c.Cli = http.DefaultClient
	}
	if c.Key == "" {
		c.Key = os.Getenv("INDEXINGCO_API_KEY")
	}
	if c.Key == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Key must not be empty", c)))
	}
	if c.Url == "" {
		c.Url = "https://app.indexing.co"
	}

	return &Client{
		cli: c.Cli,
		key: c.Key,
		url: strings.TrimSuffix(strings.TrimSpace(c.Url), "/"),
	}
}
