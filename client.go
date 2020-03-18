package gnhentai

import "io"

type Client interface {
	ByID(id int) (Doujinshi, error)
	Random() (Doujinshi, error)

	Search(q string, page int) ([]Doujinshi, error)
	SearchByTag(tag string, page int) ([]Doujinshi, error)
	Related(id int) ([]Doujinshi, error)
}

type Downloader interface {
	Page(mediaID, n int) (io.ReadCloser, error)
	Thumbnail(mediaID int, n int) (io.ReadCloser, error)
	Cover(mediaID int) (io.ReadCloser, error)
}
