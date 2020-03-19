package parser

import (
	"fmt"
	"github.com/tdakkota/gnhentai"
	"io"
	"net/http"
)

type Parser struct {
	client *http.Client
}

func NewParser(opts ...Option) *Parser {
	c := &Parser{}

	for _, opt := range opts {
		opt(c)
	}

	if c.client == nil {
		c.client = http.DefaultClient
	}

	return c
}

func (c Parser) ByID(id int) (gnhentai.Doujinshi, error) {
	return c.requestComic(fmt.Sprintf("%s/g/%d/", gnhentai.BaseNHentaiLink, id))
}

func (c Parser) Random() (gnhentai.Doujinshi, error) {
	return c.requestComic(gnhentai.BaseNHentaiLink + "/random/")
}

func (c Parser) requestComic(url string) (gnhentai.Doujinshi, error) {
	r, err := c.request(url)

	if err != nil {
		return gnhentai.Doujinshi{}, err
	}
	defer r.Close()

	return ParseComic(r)
}

func (c Parser) Search(q string, page int) ([]gnhentai.Doujinshi, error) {
	var u string // url

	if page >= 2 {
		u = fmt.Sprintf("%s/search/?q=%s&page=%d", gnhentai.BaseNHentaiLink, q, page)
	} else {
		u = fmt.Sprintf("%s/search/?q=%s", gnhentai.BaseNHentaiLink, q)
	}

	return c.requestSearch(u)
}

func (c Parser) SearchByTag(tag gnhentai.Tag, page int) ([]gnhentai.Doujinshi, error) {
	var u string // url

	if page >= 2 {
		u = fmt.Sprintf("%s/tag/%s/?page=%d", gnhentai.BaseNHentaiLink, tag.Name, page)
	} else {
		u = fmt.Sprintf("%s/tag/%s/", gnhentai.BaseNHentaiLink, tag.Name)
	}

	return c.requestSearch(u)
}

func (c Parser) Related(id int) ([]gnhentai.Doujinshi, error) {
	r, err := c.request(fmt.Sprintf("%s/g/%d/", gnhentai.BaseNHentaiLink, id))
	if err != nil {
		return nil, err
	}

	return ParseRelated(r)
}

func (c Parser) requestSearch(url string) ([]gnhentai.Doujinshi, error) {
	r, err := c.request(url)

	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ParseSearch(r)
}

func (c Parser) request(url string) (io.ReadCloser, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("bad http code: %d", r.StatusCode)
	}

	return r.Body, err
}
