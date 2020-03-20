package gnhentai

import (
	"io"
	"os"
)

type Downloader interface {
	Page(mediaID, n int) (io.ReadCloser, error)
	Thumbnail(mediaID int, n int) (io.ReadCloser, error)
	Cover(mediaID int) (io.ReadCloser, error)
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
