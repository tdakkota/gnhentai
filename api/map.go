package api

import (
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/internal/nhentaiapi"
)

func mapBook(d nhentaiapi.Book) (gnhentai.Doujinshi, error) {
	id := d.ID.String
	if n, ok := d.ID.GetInt(); ok {
		id = strconv.Itoa(n)
	}

	title, err := mapTitle(d.Title)
	if err != nil {
		return gnhentai.Doujinshi{}, errors.Wrap(err, "map title")
	}
	tags, err := mapSlice(d.Tags, mapTag)
	if err != nil {
		return gnhentai.Doujinshi{}, errors.Wrap(err, "map tags")
	}
	images, err := mapImages(d.Images)
	if err != nil {
		return gnhentai.Doujinshi{}, errors.Wrap(err, "map images")
	}
	return gnhentai.Doujinshi{
		ID:           id,
		MediaID:      d.MediaID.Or(""),
		Title:        title,
		Tags:         tags,
		Scanlator:    d.Scanlator.Or(""),
		NumPages:     d.NumPages.Or(0),
		NumFavorites: d.NumFavorites.Or(0),
		UploadDate:   d.UploadDate.Or(time.Time{}),
		Images:       images,
	}, nil
}

func mapTitle(d nhentaiapi.Title) (gnhentai.Title, error) {
	return gnhentai.Title{
		English:  d.English.Or(""),
		Japanese: d.Japanese.Or(""),
		Pretty:   d.Pretty.Or(""),
	}, nil
}

func mapTag(d nhentaiapi.Tag) (gnhentai.Tag, error) {
	return gnhentai.Tag{
		ID:    d.ID,
		Type:  string(d.Type),
		Name:  d.Name,
		URL:   d.URL.Or(""),
		Count: d.Count.Or(0),
	}, nil
}

func mapImages(d nhentaiapi.Images) (gnhentai.Images, error) {
	pages, err := mapSlice(d.Pages, mapImage)
	if err != nil {
		return gnhentai.Images{}, errors.Wrap(err, "map pages")
	}
	cover, err := mapImage(d.Cover.Or(nhentaiapi.Image{}))
	if err != nil {
		return gnhentai.Images{}, errors.Wrap(err, "map cover")
	}
	thumbnail, err := mapImage(d.Thumbnail.Or(nhentaiapi.Image{}))
	if err != nil {
		return gnhentai.Images{}, errors.Wrap(err, "map thumbnail")
	}
	return gnhentai.Images{
		Pages:     pages,
		Cover:     cover,
		Thumbnail: thumbnail,
	}, nil
}

func mapImage(d nhentaiapi.Image) (gnhentai.Image, error) {
	return gnhentai.Image{
		T:      d.T,
		Width:  d.W.Or(0),
		Height: d.H.Or(0),
	}, nil
}

func mapSlice[From, To any](from []From, mapper func(From) (To, error)) (to []To, _ error) {
	to = make([]To, len(from))
	for i, v := range from {
		m, err := mapper(v)
		if err != nil {
			return to, err
		}
		to[i] = m
	}
	return to, nil
}
