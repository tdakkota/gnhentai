package gnhentai

import (
	"fmt"
	"io"
	"os"
)

// Provides interface to download doujinshi images
type Downloader interface {
	// Returns page reader
	Page(mediaID, n int, format string) (io.ReadCloser, error)
	/// Returns page thumbnail reader
	Thumbnail(mediaID int, n int, format string) (io.ReadCloser, error)
	// Returns cover reader
	Cover(mediaID int, format string) (io.ReadCloser, error)
}

// Returns link to doujinshi page
func PageLink(mediaID, n int, format string) string {
	return fmt.Sprintf("https://i.nhentai.net/galleries/%d/%d.%s", mediaID, n, format)
}

// Returns link to doujinshi page thumbnail
func ThumbnailLink(mediaID, n int, format string) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/%dt.%s", mediaID, n, format)
}

// Returns link to doujinshi cover
func CoverLink(mediaID int, format string) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/cover.%s", mediaID, format)
}

// Returns image extension from Image object
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

/// Downloads all pages of passed Doujinshi, using downloader
/// The filename is specified by namer callback
func DownloadAll(downloader Downloader, d Doujinshi, namer func(i int, d Doujinshi) string) error {
	for i := 0; i < d.NumPages; i++ {
		name := namer(i+1, d)
		if err := downloadOne(downloader, d, i+1, name); err != nil {
			return err
		}
	}
	return nil
}
