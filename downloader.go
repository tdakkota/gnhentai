package gnhentai

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ThumbnailLink(mediaID, n int) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/%dt.jpg", mediaID, n)
}

func CoverLink(mediaID int) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/cover.jpg", mediaID)
}

func ImageLink(mediaID, n int) string {
	return fmt.Sprintf("https://i.nhentai.net/galleries/%d/%d.jpg", mediaID, n)
}

func PageLink(ID, n int) string {
	return fmt.Sprintf("https://nhentai.net/g/%d/%d/", ID, n)
}

func MainPageLink(ID int) string {
	return fmt.Sprintf("https://nhentai.net/g/%d", ID)
}

func downloadOne(downloader Downloader, mediaID, i int, name string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	r, err := downloader.Page(mediaID, i+1)
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = io.Copy(file, r)
	return err
}

func DownloadAll(downloader Downloader, d Doujinshi, namer func(i int, d Doujinshi) string) error {
	for i := 0; i < d.NumPages; i++ {
		name := namer(i+1, d)
		if err := downloadOne(downloader, d.MediaID, i+1, name); err != nil {
			return err
		}
	}
	return nil
}

type SimpleDownloader struct {
	client *http.Client
}

func NewSimpleDownloader(client *http.Client) SimpleDownloader {
	if client == nil {
		client = http.DefaultClient
	}
	return SimpleDownloader{client: client}
}

func (s SimpleDownloader) download(url string) (io.ReadCloser, error) {
	r, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

func (s SimpleDownloader) Page(mediaID, n int) (io.ReadCloser, error) {
	return s.download(ImageLink(mediaID, n))
}

func (s SimpleDownloader) Cover(mediaID int) (io.ReadCloser, error) {
	return s.download(CoverLink(mediaID))
}

func (s SimpleDownloader) Thumbnail(mediaID, n int) (io.ReadCloser, error) {
	return s.download(ThumbnailLink(mediaID, n))
}
