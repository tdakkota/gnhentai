package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tdakkota/gnhentai"
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
	if r.Body != nil {
		defer r.Body.Close()
	}

	u := r.Request.URL.String()
	_, err = fmt.Sscanf(u, "https://nhentai.net/g/%d/", &id)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID in %s: %v", u, err)
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
	Result   []gnhentai.Doujinshi `json:"result"`
	NumPages int                  `json:"num_pages"`
	PerPage  int                  `json:"per_page"`
	Error    string               `json:"error"`
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

func (c Client) Page(mediaID, n int, format string) (io.ReadCloser, error) {
	return c.request(gnhentai.PageLink(mediaID, n, format))
}

func (c Client) Thumbnail(mediaID int, n int, format string) (io.ReadCloser, error) {
	return c.request(gnhentai.ThumbnailLink(mediaID, n, format))
}

func (c Client) Cover(mediaID int, format string) (io.ReadCloser, error) {
	return c.request(gnhentai.CoverLink(mediaID, format))
}

func (c Client) request(url string) (io.ReadCloser, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		if r.Body != nil {
			_ = r.Body.Close()
		}
		return nil, fmt.Errorf("bad http code: %d", r.StatusCode)
	}

	return r.Body, err
}
