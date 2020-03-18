package gnhentai

import (
	"fmt"
	"time"
)

const BaseNHentaiLink = "https://nhentai.net"

type Tag struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}

type Image struct {
	T      string `json:"t"`
	Width  int    `json:"w"`
	Height int    `json:"h"`
}

type Page = Image
type Cover = Image
type Thumbnail = Image

type Images struct {
	Pages     []Page    `json:"pages"`
	Cover     Cover     `json:"cover"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

type Title struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Pretty   string `json:"pretty"`
}

type Doujinshi struct {
	// ID is unique identification number of Doujinshi.
	// Note: parser does not parse ID of Doujinshi.
	ID        int    `json:"id"`
	MediaID   int    `json:"media_id"`
	Title     Title  `json:"name"`
	Tags      []Tag  `json:"tags"`
	Scanlator string `json:"scanlator"`
	// NumPages is number of manga pages.
	NumPages     int       `json:"num_pages"`
	NumFavorites int       `json:"num_favorites"`
	UploadDate   time.Time `json:"upload_date"`
	Images       Images    `json:"images"`
}

func (d Doujinshi) Thumbnail(n int) string {
	return ThumbnailLink(d.MediaID, n)
}

func (d Doujinshi) Page(n int) string {
	return PageLink(d.MediaID, n)
}

func ThumbnailLink(mediaID, n int) string {
	return fmt.Sprintf("https://t.nhentai.net/galleries/%d/%dt.jpg", mediaID, n)
}

func PageLink(mediaID, n int) string {
	return fmt.Sprintf("https://i.nhentai.net/galleries/%d/%d.jpg", mediaID, n)
}
