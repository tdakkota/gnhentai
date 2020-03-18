package parser

import (
	"net/http"
)

type Option func(c *Parser)

func WithClient(client *http.Client) Option {
	return func(c *Parser) {
		c.client = client
	}
}
