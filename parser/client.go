package parser

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-faster/errors"
	"github.com/tdakkota/gnhentai"
)

// Parser is a simple scraper for nhentai.net.
type Parser struct {
	baseURL *url.URL
	client  *http.Client
}

// ParserOptions defines options for [Parser].
type ParserOptions struct {
	// Client is a http client to use for requests.
	Client *http.Client
}

func (opts *ParserOptions) setDefaults() {
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
}

var defaultBaseURL = errors.Must(url.Parse(gnhentai.BaseNHentaiLink))

// NewParser creates a new [Parser] with options.
func NewParser(opts ParserOptions) *Parser {
	opts.setDefaults()

	return &Parser{
		baseURL: defaultBaseURL,
		client:  opts.Client,
	}
}

// ByID returns book metadata by ID.
func (c *Parser) ByID(ctx context.Context, id int) (gnhentai.Doujinshi, error) {
	return c.requestComic(ctx, c.baseURL.JoinPath("g", strconv.Itoa(id), "/"))
}

// Random returns random book metadata.
func (c *Parser) Random(ctx context.Context) (gnhentai.Doujinshi, error) {
	return c.requestComic(ctx, c.baseURL.JoinPath("random/"))
}

// Search books by term.
func (c *Parser) Search(ctx context.Context, q string, page int) ([]gnhentai.Doujinshi, error) {
	u := c.baseURL.JoinPath("search/")

	query := u.Query()
	query.Set("q", q)
	if page >= 2 {
		query.Set("page", strconv.Itoa(page))
	}
	u.RawQuery = query.Encode()

	return c.requestSearch(ctx, u)
}

// SearchByTag searches books by given [gnhentai.Tag].
func (c *Parser) SearchByTag(ctx context.Context, tag gnhentai.Tag, page int) ([]gnhentai.Doujinshi, error) {
	u := c.baseURL.JoinPath("tag", tag.Name, "/")

	if page >= 2 {
		query := u.Query()
		query.Set("page", strconv.Itoa(page))
		u.RawQuery = query.Encode()
	}

	return c.requestSearch(ctx, u)
}

// Related returns related books.
func (c *Parser) Related(ctx context.Context, id int) ([]gnhentai.Doujinshi, error) {
	doc, err := c.scrapeHTML(ctx, c.baseURL.JoinPath("g", strconv.Itoa(id), "/"))
	if err != nil {
		return nil, err
	}
	return ParseRelated(doc.Selection)
}

func (c *Parser) requestComic(ctx context.Context, u *url.URL) (gnhentai.Doujinshi, error) {
	doc, err := c.scrapeHTML(ctx, u)
	if err != nil {
		return gnhentai.Doujinshi{}, err
	}
	return ParseComic(doc.Selection)
}

func (c *Parser) requestSearch(ctx context.Context, u *url.URL) ([]gnhentai.Doujinshi, error) {
	doc, err := c.scrapeHTML(ctx, u)
	if err != nil {
		return nil, err
	}
	return ParseSearch(doc.Selection)
}

func (c *Parser) scrapeHTML(ctx context.Context, u *url.URL) (*goquery.Document, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "make request")
	}

	r, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send request")
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, errors.Errorf("bad http code: %d", r.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "parse html")
	}

	return doc, nil
}
