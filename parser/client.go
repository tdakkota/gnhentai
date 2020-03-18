package parser

import (
	"fmt"
	"github.com/tdakkota/gnhentai"
	"net/http"
)

type Parser struct {
	client *http.Client
}

func NewClient(opts ...Option) *Parser {
	c := &Parser{}

	for _, opt := range opts {
		opt(c)
	}

	if c.client == nil {
		c.client = http.DefaultClient
	}

	return c
}

func (c *Parser) ByID(id int) (gnhentai.Doujinshi, error) {
	return c.request(fmt.Sprintf("%s/g/%d/", gnhentai.BaseNHentaiLink, id))
}

func (c *Parser) Random() (gnhentai.Doujinshi, error) {
	return c.request(gnhentai.BaseNHentaiLink + "/random/")
}

func (c *Parser) request(url string) (gnhentai.Doujinshi, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return gnhentai.Doujinshi{}, err
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if r.StatusCode != 200 {
		return gnhentai.Doujinshi{}, fmt.Errorf("bad http code: %d", r.StatusCode)
	}

	return Parse(r.Body)
}

func (c *Parser) Search(q string) ([]gnhentai.Doujinshi, error) {
	panic("implement me!")
}

func (c *Parser) SearchByTag(tag gnhentai.Tag) ([]gnhentai.Doujinshi, error) {
	panic("implement me!")
}

func (c *Parser) Related(d gnhentai.Doujinshi) ([]gnhentai.Doujinshi, error) {
	panic("implement me!")
}
