package api

import (
	"github.com/tdakkota/gnhentai"
	"io"
)

type Cache interface {
	GetDoujinshi(bookID int) (gnhentai.Doujinshi, error)
	SetDoujinshi(bookID int, d gnhentai.Doujinshi) error

	GetPage(bookID, n int) (io.ReadCloser, error)
	SetPage(bookID int, image io.ReadCloser) error

	GetCover(bookID int) (io.ReadCloser, error)
	SetCover(bookID int, image io.ReadCloser) error

	GetThumbnail(bookID int) (io.ReadCloser, error)
	SetThumbnail(bookID int, image io.ReadCloser) error

	Search(q string) ([]gnhentai.Doujinshi, error)
	SetSearch(q string, result []gnhentai.Doujinshi) error

	SearchByTag(tag gnhentai.Tag) ([]gnhentai.Doujinshi, error)
	SetSearchByTag(tag gnhentai.Tag, result []gnhentai.Doujinshi) error

	Related(bookID int) ([]gnhentai.Doujinshi, error)
	SetRelated(bookID int, result []gnhentai.Doujinshi) error
}
