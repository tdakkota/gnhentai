package api

import (
	"net/http"
)

type Option func(c *Client)

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}
