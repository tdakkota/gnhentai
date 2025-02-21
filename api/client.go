package api

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/internal/nhentaiapi"
)

// BaseNHentaiAPILink defines base link for nhentai.net API.
const BaseNHentaiAPILink = gnhentai.BaseNHentaiLink + "/api"

// Client is a nhentai.net API client.
type Client struct {
	client    *http.Client
	apiclient *nhentaiapi.Client
}

var _ gnhentai.Client = (*Client)(nil)

// ClientOptions defines options for nhentai.net API client.
type ClientOptions struct {
	// Client sets HTTP client to use.
	Client *http.Client
}

func (o ClientOptions) setDefaults() {
	if o.Client != nil {
		o.Client = http.DefaultClient
	}
}

var baseURL = errors.Must(url.Parse(BaseNHentaiAPILink))

// NewClient creates a new nhentai.net API client.
func NewClient(opts ClientOptions) *Client {
	opts.setDefaults()

	apiclient, err := nhentaiapi.NewClient(baseURL.String(), nhentaiapi.WithClient(opts.Client))
	if err != nil {
		panic(err)
	}

	return &Client{
		client:    opts.Client,
		apiclient: apiclient,
	}
}

// ByID returns book metadata by ID.
func (c *Client) ByID(ctx context.Context, id int) (d gnhentai.Doujinshi, _ error) {
	r, err := c.apiclient.GetBook(ctx, nhentaiapi.GetBookParams{BookID: strconv.Itoa(id)})
	if err != nil {
		return d, c.mapError(err)
	}
	d, err = mapBook(*r)
	if err != nil {
		return d, errors.Wrap(err, "map book")
	}
	return d, nil
}

// Random returns random book metadata.
func (c *Client) Random(ctx context.Context) (d gnhentai.Doujinshi, _ error) {
	id, err := c.randomPage(ctx)
	if err != nil {
		return d, errors.Wrap(err, "get random page")
	}

	return c.ByID(ctx, id)
}

func (c *Client) randomPage(ctx context.Context) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, gnhentai.BaseNHentaiLink+"/random/", nil)
	if err != nil {
		return 0, errors.Wrap(err, "create request")
	}

	r, err := c.client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "do requests")
	}
	defer func() {
		_ = r.Body.Close()
	}()

	redirect := r.Request.URL
	id, err := strconv.Atoi(path.Base(redirect.Path))
	if err != nil {
		return 0, errors.Errorf("invalid random page redirect: %q", redirect)
	}
	return id, nil
}

// Search books by term.
func (c *Client) Related(ctx context.Context, id int) ([]gnhentai.Doujinshi, error) {
	r, err := c.apiclient.Related(ctx, nhentaiapi.RelatedParams{BookID: strconv.Itoa(id)})
	if err != nil {
		return nil, c.mapError(err)
	}
	d, err := mapSlice(r.Result, mapBook)
	if err != nil {
		return nil, errors.Wrap(err, "map books")
	}
	return d, nil
}

// SearchByTag searches books by given [gnhentai.Tag].
func (c *Client) Search(ctx context.Context, q string, page int) ([]gnhentai.Doujinshi, error) {
	r, err := c.apiclient.Search(ctx, nhentaiapi.SearchParams{
		Query: q,
		Page:  nhentaiapi.NewOptInt(page),
	})
	if err != nil {
		return nil, c.mapError(err)
	}
	d, err := mapSlice(r.Result, mapBook)
	if err != nil {
		return nil, errors.Wrap(err, "map books")
	}
	return d, nil
}

// Related returns related books.
func (c *Client) SearchByTag(ctx context.Context, tag gnhentai.Tag, page int) ([]gnhentai.Doujinshi, error) {
	r, err := c.apiclient.SearchByTagID(ctx, nhentaiapi.SearchByTagIDParams{
		TagID: tag.ID,
		Page:  nhentaiapi.NewOptInt(page),
	})
	if err != nil {
		return nil, c.mapError(err)
	}
	d, err := mapSlice(r.Result, mapBook)
	if err != nil {
		return nil, errors.Wrap(err, "map books")
	}
	return d, nil
}

func (c *Client) mapError(err error) error {
	if apiErr, ok := errors.Into[*nhentaiapi.ErrorStatusCode](err); ok {
		return &APIError{
			StatusCode: apiErr.StatusCode,
			Message:    apiErr.Response.Error,
		}
	}
	return err
}
