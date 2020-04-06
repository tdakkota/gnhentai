package downloaders

import (
	"github.com/tdakkota/gnhentai"
	"io"
	"net/http"
)

type SimpleDownloader struct {
	client *http.Client
}

func NewSimpleDownloader(client *http.Client) SimpleDownloader {
	if client == nil {
		client = http.DefaultClient
	}
	return SimpleDownloader{client: client}
}

func (d SimpleDownloader) download(url string) (io.ReadCloser, error) {
	r, err := d.client.Get(url)
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

func (d SimpleDownloader) Page(mediaID, n int, format string) (io.ReadCloser, error) {
	return d.download(gnhentai.PageLink(mediaID, n, format))
}

func (d SimpleDownloader) Cover(mediaID int, format string) (io.ReadCloser, error) {
	return d.download(gnhentai.CoverLink(mediaID, format))
}

func (d SimpleDownloader) Thumbnail(mediaID int, n int, format string) (io.ReadCloser, error) {
	return d.download(gnhentai.ThumbnailLink(mediaID, n, format))
}
