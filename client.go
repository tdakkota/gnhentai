package gnhentai

type Client interface {
	ByID(id int) (Doujinshi, error)
	Random() (Doujinshi, error)

	Search(q string, page int) ([]Doujinshi, error)
	SearchByTag(tag Tag, page int) ([]Doujinshi, error)
	Related(id int) ([]Doujinshi, error)
}
