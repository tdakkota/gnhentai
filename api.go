package gnhentai

import (
	"context"
	"net/http"
	"path"

	"github.com/go-faster/errors"

	"github.com/tdakkota/gnhentai/nhentaiapi"
)

// API provides API-based [Client] implementation.
type API struct {
	client *http.Client
	api    nhentaiapi.Invoker
}

var _ Client = (*API)(nil)

// APIOptions defines options for nhentai.net API client.
type APIOptions struct {
	// Client sets HTTP client to use.
	Client *http.Client
	// API sets API client to use.
	API nhentaiapi.Invoker
}

func (o *APIOptions) setDefaults() {
	if o.Client == nil {
		o.Client = http.DefaultClient
	}
	if o.API == nil {
		o.API = errors.Must(nhentaiapi.NewClient(
			nhentaiapi.BaseNHentaiLink.String(),
			nhentaiapi.WithClient(o.Client),
		))
	}
}

// NewAPI creates a new nhentai.net API client.
func NewAPI(opts APIOptions) *API {
	opts.setDefaults()

	return &API{
		client: opts.Client,
		api:    opts.API,
	}
}

// ByID returns book metadata by ID.
func (c *API) ByID(ctx context.Context, id BookID) (d *nhentaiapi.Book, _ error) {
	r, err := c.api.GetBook(ctx, nhentaiapi.GetBookParams{BookID: id})
	if err != nil {
		return nil, c.mapError(err)
	}
	return r, nil
}

// Random returns random book metadata.
func (c *API) Random(ctx context.Context) (d *nhentaiapi.Book, _ error) {
	id, err := c.randomPage(ctx)
	if err != nil {
		return d, errors.Wrap(err, "get random page")
	}
	return c.ByID(ctx, id)
}

var randomPage = nhentaiapi.BaseNHentaiLink.JoinPath("/random/").String()

func (c *API) randomPage(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, randomPage, nil)
	if err != nil {
		return "", errors.Wrap(err, "create request")
	}

	r, err := c.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "do request")
	}
	defer func() {
		_ = r.Body.Close()
	}()

	redirect := r.Request.URL
	id := path.Base(redirect.Path)
	if id == "" {
		return "", errors.Errorf("invalid redirect URL: %q", redirect)
	}
	return id, nil
}

// Search books by term.
func (c *API) Search(ctx context.Context, q string, page int) (*nhentaiapi.SearchResponse, error) {
	r, err := c.api.Search(ctx, nhentaiapi.SearchParams{
		Query:   q,
		Page:    nhentaiapi.NewOptInt(page),
		PerPage: 25,
	})
	if err != nil {
		return nil, c.mapError(err)
	}
	return r, nil
}

// SearchByTag searches books by given [nhentaiapi.Tag].
func (c *API) SearchByTag(ctx context.Context, tag nhentaiapi.Tag, page int) (*nhentaiapi.SearchResponse, error) {
	r, err := c.api.SearchByTagID(ctx, nhentaiapi.SearchByTagIDParams{
		TagID:   tag.ID,
		Page:    nhentaiapi.NewOptInt(page),
		PerPage: 25,
	})
	if err != nil {
		return nil, c.mapError(err)
	}
	return r, nil
}

func (c *API) mapError(err error) error {
	return err
}

// Related returns related books.
func (c *API) Related(ctx context.Context, id BookID) (*nhentaiapi.SearchResponse, error) {
	r, err := c.api.Related(ctx, nhentaiapi.RelatedParams{BookID: id})
	if err != nil {
		return nil, c.mapError(err)
	}
	return r, nil
}
