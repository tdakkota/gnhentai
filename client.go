package gnhentai

import (
	"context"

	"github.com/tdakkota/gnhentai/nhentaiapi"
)

type (
	// BookID defines book ID.
	BookID = string
)

// Client is an abstraction for different nhentai.net API implementations
type Client interface {
	// ByID returns book metadata by ID.
	ByID(ctx context.Context, id BookID) (*nhentaiapi.Book, error)

	// Random returns random book metadata.
	Random(ctx context.Context) (*nhentaiapi.Book, error)

	// Search books by term
	//
	// 	- You can search for multiple terms at the same time, and this will return only galleries that contain both terms.
	//    For example,
	//
	//    	anal tanlines
	//
	//    finds all galleries that contain both 'anal' and 'tanlines'.
	//
	//  - You can exclude terms by prefixing them with '-'.
	//    For example,
	//
	//    	anal tanlines -yaoi
	//
	//    matches all galleries matching 'anal' and 'tanlines' but not 'yaoi'.
	//
	//  - Exact searches can be performed by wrapping terms in double quotes.
	//    For example, "big breasts" only matches galleries with "big breasts" exactly somewhere in the title or in tags.
	//
	//  - These can be combined with tag namespaces for finer control over the query: parodies:railgun -tag:\"big breasts\".
	Search(ctx context.Context, q string, page int) (*nhentaiapi.SearchResponse, error)

	// SearchByTag searches books by given [Tag].
	// Note: API client uses [Tag.ID], when Parser uses [Tag.Name] to get metadata, so both fields should not be empty.
	SearchByTag(ctx context.Context, tag nhentaiapi.Tag, page int) (*nhentaiapi.SearchResponse, error)
}
