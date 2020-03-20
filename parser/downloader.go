package parser

import (
	"fmt"
	"io"
	"net/http"
)

func ThumbnailLink(mediaID, n int, format string) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/%dt.%s", mediaID, n, format)
}

func CoverLink(mediaID int, format string) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/cover.%s", mediaID, format)
}

func ImageLink(mediaID, n int, format string) string {
	return fmt.Sprintf("https://i.nhentai.net/galleries/%d/%d.%s", mediaID, n, format)
}

type Downloader struct {
	client *http.Client
	format string
}

func NewDownloader(client *http.Client) Downloader {
	if client == nil {
		client = http.DefaultClient
	}
	return Downloader{client: client, format: "jpg"}
}

func (d *Downloader) ImageFormat() string {
	return d.format
}

func (d *Downloader) SetImageFormat(format string) {
	d.format = format
}

func (d Downloader) download(url string) (io.ReadCloser, error) {
	r, err := d.client.Get(url)
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

func (d Downloader) Page(mediaID, n int) (io.ReadCloser, error) {
	return d.download(ImageLink(mediaID, n, d.format))
}

func (d Downloader) Cover(mediaID int) (io.ReadCloser, error) {
	return d.download(CoverLink(mediaID, d.format))
}

func (d Downloader) Thumbnail(mediaID, n int) (io.ReadCloser, error) {
	return d.download(ThumbnailLink(mediaID, n, d.format))
}
