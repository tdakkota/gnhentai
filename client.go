package gnhentai

import (
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(opts ...Option) *Client {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if c.client == nil {
		c.client = http.DefaultClient
	}

	return c
}

func (c *Client) ByID(id int) (Doujinshi, error) {
	return c.request(fmt.Sprintf("%s/g/%d/", BaseNHentaiLink, id))
}

func (c *Client) Random() (Doujinshi, error) {
	return c.request(BaseNHentaiLink + "/random/")
}

func (c *Client) request(url string) (Doujinshi, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return Doujinshi{}, err
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if r.StatusCode != 200 {
		return Doujinshi{}, fmt.Errorf("bad http code: %d", r.StatusCode)
	}

	return Parse(r.Body)
}
