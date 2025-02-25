package nhentaiapi

import (
	"fmt"
	"math/rand/v2"
	"net/url"
	"strconv"

	"github.com/go-faster/errors"
)

// BaseNHentaiLink is default URL of nhentai.net.
var BaseNHentaiLink = errors.Must(url.Parse("https://nhentai.net"))

// Ext returns image extension.
func (i Image) Ext() (string, error) {
	switch i.T {
	case "j":
		return "jpg", nil
	case "p":
		return "png", nil
	case "w":
		return "webp.webp", nil
	default:
		return "", errors.Errorf("unknown image type %q", i.T)
	}
}

// ToString returns string representation of BookID.
func (id *BookID) ToString() string {
	if v, ok := id.GetInt(); ok {
		return strconv.Itoa(v)
	}
	return id.String
}

// Name returns pretty name of Doujinshi.
func (d Book) Name() string {
	for _, r := range []OptNilString{
		d.Title.Pretty,
		d.Title.English,
		d.Title.Japanese,
	} {
		if v, ok := r.Get(); ok && v != "" {
			return v
		}
	}
	return ""
}

func (d Book) getPageExt(page int) (string, error) {
	if page < 1 {
		return "", errors.Errorf("page %d is invalid", page)
	}
	page--

	pages := d.Images.Pages
	if len(pages) <= page {
		return "", errors.Errorf("page %d of %d not found", page, len(pages))
	}

	return pages[page].Ext()
}

// PageLink returns link to doujinshi page.
func (d Book) PageLink(page int) (string, error) {
	ext, err := d.getPageExt(page)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s/galleries/%s/%d.%s", randomCDN(), d.MediaID, page, ext), nil
}

// ThumbnailLink returns link to doujinshi page thumbnail.
func (d Book) ThumbnailLink(page int) (string, error) {
	ext, err := d.getPageExt(page)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s/galleries/%s/%dt.%s", randomCDN(), d.MediaID, page, ext), nil
}

// CoverLink returns link to doujinshi cover.
func (d Book) CoverLink() (string, error) {
	cover, ok := d.Images.Cover.Get()
	if !ok {
		return "", errors.New("book does not have cover")
	}
	ext, err := cover.Ext()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s/galleries/%s/cover.%s", randomCDN(), d.MediaID, ext), nil
}

func randomCDN() string {
	return fmt.Sprintf("t%d.nhentai.net", rand.N(4)+1)
}
