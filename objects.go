package gnhentai

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"time"
)

const BaseNHentaiLink = "https://nhentai.net"

type BaseTag struct {
	ID    int
	Count int
	Name  string
	Link  string
}

type Parody = BaseTag
type Character = BaseTag
type Tag = BaseTag
type Artist = BaseTag
type Group = BaseTag
type Category = BaseTag
type Language = BaseTag

type Image struct {
	Link   string
	Width  int
	Height int
}

type Preview struct {
	Number    int
	Link      string
	Thumbnail Image
}

type Previews []Preview

type Doujinshi struct {
	// ID is unique identification number of Doujinshi.
	// Note: parser does not parse ID of Doujinshi.
	ID      int
	MediaID int
	// Name is pretty english name
	Name string
	// AlterName is original name
	AlterName string

	// Tags.
	Parodies   []Parody
	Characters []Character
	Tags       []Tag
	Artists    []Artist
	Groups     []Group
	Languages  []Language
	Categories []Category

	// Previews is slice of pages previews.
	Previews Previews

	// Length is number of manga pages.
	Length   int
	Uploaded time.Time
}

func Parse(r io.Reader) (result Doujinshi, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return
	}
	return parse(doc)
}
