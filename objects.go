package gnhentai

import (
	"encoding/json"
	"fmt"
	"strconv"
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

// JSONTimestamp
type JSONTimestamp struct {
	time.Time
}

// UnmarshalJSON parses json number into time.Time
func (t *JSONTimestamp) UnmarshalJSON(b []byte) error {
	parsed, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(parsed, 0)
	return nil
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

func (d *Doujinshi) UnmarshalJSON(data []byte) error {
	type Alias Doujinshi
	correctStruct := struct {
		*Alias
		MediaID    string        `json:"media_id"`
		UploadDate JSONTimestamp `json:"upload_date"`
	}{}

	err := json.Unmarshal(data, &correctStruct)
	if err != nil {
		return err
	}

	d.ID = correctStruct.ID
	d.MediaID, err = strconv.Atoi(correctStruct.MediaID)
	if err != nil {
		return fmt.Errorf("failed to parse media id from '%s': %v", correctStruct.MediaID, err)
	}

	d.Title = correctStruct.Title
	d.Tags = correctStruct.Tags
	d.Scanlator = correctStruct.Scanlator
	d.NumPages = correctStruct.NumPages
	d.NumFavorites = correctStruct.NumFavorites
	d.UploadDate = correctStruct.UploadDate.Time
	d.Images = correctStruct.Images

	return nil
}
