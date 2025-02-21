package gnhentai

import "context"

// Client is an abstraction for different nhentai.net API implementations
type Client interface {
	// ByID returns book metadata by ID.
	ByID(ctx context.Context, id int) (Doujinshi, error)

	// Random returns random book metadata.
	Random(ctx context.Context) (Doujinshi, error)

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
	Search(ctx context.Context, q string, page int) ([]Doujinshi, error)

	// SearchByTag searches books by given [Tag].
	// Note: API client uses [Tag.ID], when Parser uses [Tag.Name] to get metadata, so both fields should not be empty.
	SearchByTag(ctx context.Context, tag Tag, page int) ([]Doujinshi, error)

	// Related returns related books.
	Related(ctx context.Context, id int) ([]Doujinshi, error)
}
