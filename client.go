package gnhentai

import "io"

type Client interface {
	ByID(id int) (Doujinshi, error)
	Random() (Doujinshi, error)

	Search(q string) ([]Doujinshi, error)
	SearchByTag(tag Tag) ([]Doujinshi, error)
	Related(d Doujinshi) ([]Doujinshi, error)
}

type Downloader interface {
	Page(doujinshiID, n int) (io.ReadCloser, error)
	Cover(doujinshiID int) (io.ReadCloser, error)
	Thumbnail(doujinshiID int) (io.ReadCloser, error)
}
