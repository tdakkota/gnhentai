package gnhentai

import (
	"io"
	"os"
)

type Downloader interface {
	Page(mediaID, n int, format string) (io.ReadCloser, error)
	Thumbnail(mediaID int, n int, format string) (io.ReadCloser, error)
	Cover(mediaID int, format string) (io.ReadCloser, error)
}

func FormatFromImage(image Image) string {
	switch image.T {
	case "j":
		return "jpg"
	case "p":
		return "png"
	}

	return ""
}

func downloadOne(downloader Downloader, d Doujinshi, i int, name string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	format := FormatFromImage(d.Images.Pages[i])
	if format == "" {
		format = "jpg"
	}

	r, err := downloader.Page(d.MediaID, i+1, format)
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
		if err := downloadOne(downloader, d, i+1, name); err != nil {
			return err
		}
	}
	return nil
}
