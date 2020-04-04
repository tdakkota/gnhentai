package gnhentai

// Client is an abstraction for different nhentai.net API implementations
type Client interface {
	// Returns book metadata by ID
	ByID(id int) (Doujinshi, error)

	// Returns random book metadata
	Random() (Doujinshi, error)

	// Search books by term
	// * You can search for multiple terms at the same time, and this will return only galleries that contain both terms.
	//   For example, anal tanlines finds all galleries that contain both anal and tanlines.
	// * You can exclude terms by prefixing them with -. For example, anal tanlines -yaoi matches all galleries matching anal and tanlines but not yaoi.
	// * Exact searches can be performed by wrapping terms in double quotes. F
	//   For example, \"big breasts\" only matches galleries with \"big breasts\" somewhere in the title or in tags.
	// * These can be combined with tag namespaces for finer control over the query: parodies:railgun -tag:\"big breasts\".
	Search(q string, page int) ([]Doujinshi, error)

	// Search books by given Tag
	// Note: API client uses Tag.ID, when Parser uses Tag.Name to get metadata, so both fields should not be empty
	SearchByTag(tag Tag, page int) ([]Doujinshi, error)

	// Search related books
	Related(id int) ([]Doujinshi, error)
}
