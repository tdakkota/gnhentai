package api

import (
	"encoding/json"
	"fmt"
	"github.com/tdakkota/gnhentai"
	"io"
	"net/http"
)

const BaseNHentaiAPILink = gnhentai.BaseNHentaiLink + "/api"

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

type Result struct {
	gnhentai.Doujinshi
	Error string `json:"error"`
}

func (c Client) ByID(id int) (gnhentai.Doujinshi, error) {
	return c.requestComic(fmt.Sprintf("%s/gallery/%d", BaseNHentaiAPILink, id))
}

func (c Client) randomPage() (id int, err error) {
	r, err := c.client.Get(gnhentai.BaseNHentaiLink + "/random/")
	if err != nil {
		return
	}

	u, err := r.Location() // /random/ should redirect to random page
	if err != nil {
		return
	}

	_, err = fmt.Sscanf(u.String(), "https://nhentai.net/g/%d/", &id)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID in %s: %v", u.String(), err)
	}

	return
}

func (c Client) Random() (gnhentai.Doujinshi, error) {
	id, err := c.randomPage()
	if err != nil {
		return gnhentai.Doujinshi{}, fmt.Errorf("failed to get random doujinshi: %v", err)
	}
	return c.ByID(id)
}

func (c Client) requestComic(url string) (d gnhentai.Doujinshi, err error) {
	var result Result

	body, err := c.request(url)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&result)
	if err != nil {
		return
	}

	if result.Error != "" {
		err = fmt.Errorf("api error: %s", result.Error)
		return
	}

	return result.Doujinshi, nil
}

type MultipleResult struct {
	Result []gnhentai.Doujinshi `json:"result"`
	Error  string               `json:"error"`
}

func (c Client) Search(q string, page int) ([]gnhentai.Doujinshi, error) {
	var u string // url

	if page >= 2 {
		u = fmt.Sprintf("%s/galleries/search?query=%s&page=%d", BaseNHentaiAPILink, q, page)
	} else {
		u = fmt.Sprintf("%s/galleries/search?query=%s", BaseNHentaiAPILink, q)
	}

	return c.requestSearch(u)
}

func (c Client) SearchByTag(tag gnhentai.Tag, page int) ([]gnhentai.Doujinshi, error) {
	var u string // url

	if page >= 2 {
		u = fmt.Sprintf("%s/galleries/tagged?tag_id=%d&page=%d", BaseNHentaiAPILink, tag.ID, page)
	} else {
		u = fmt.Sprintf("%s/galleries/tagged?tag_id=%d", BaseNHentaiAPILink, tag.ID)
	}

	return c.requestSearch(u)
}

func (c Client) Related(id int) ([]gnhentai.Doujinshi, error) {
	return c.requestSearch(fmt.Sprintf("%s/gallery/%d/related", BaseNHentaiAPILink, id))
}

func (c Client) requestSearch(url string) ([]gnhentai.Doujinshi, error) {
	var result MultipleResult

	body, err := c.request(url)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, fmt.Errorf("api error: %s", result.Error)
	}

	return result.Result, nil
}

func (c Client) request(url string) (io.ReadCloser, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("bad http code: %d", r.StatusCode)
	}

	return r.Body, err
}
