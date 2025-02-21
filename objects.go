package gnhentai

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const BaseNHentaiLink = "https://nhentai.net"

// Tag is a structure that represents a book tag.
type Tag struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}

// Image is a structure that represents an image.
type Image struct {
	T      string `json:"t"`
	Width  int    `json:"w"`
	Height int    `json:"h"`
}

type (
	Page      = Image
	Cover     = Image
	Thumbnail = Image
)

// Images is a structure that represents a book images.
type Images struct {
	Pages     []Page    `json:"pages"`
	Cover     Cover     `json:"cover"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

// Title is a structure that represents a book title.
type Title struct {
	// English name of book
	English string `json:"english"`
	// Japanese name of book
	Japanese string `json:"japanese"`
	// Pretty(does not contain some characters) english name of book
	Pretty string `json:"pretty"`
}

// JSONTimestamp is a time.Time wrapper for json.Unmarshal
type JSONTimestamp struct {
	time.Time
}

// UnmarshalJSON implements [json.Unmarshaler].
func (t *JSONTimestamp) UnmarshalJSON(b []byte) error {
	parsed, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(parsed, 0)
	return nil
}

// Doujinshi is a structure that represents a book.
type Doujinshi struct {
	// ID is unique identification number of Doujinshi.
	// Note: parser does not parse ID of Doujinshi.
	ID int `json:"id"`
	// MediaID is unique identification number of Doujinshi images.
	MediaID int `json:"media_id"`
	// Title structure
	Title     Title  `json:"title"`
	Tags      []Tag  `json:"tags"`
	Scanlator string `json:"scanlator"`
	// NumPages is number of manga pages.
	NumPages     int       `json:"num_pages"`
	NumFavorites int       `json:"num_favorites"`
	UploadDate   time.Time `json:"upload_date"`

	Images Images `json:"images"`
}

// Name returns pretty name of Doujinshi.
func (d Doujinshi) Name() string {
	switch {
	case d.Title.Pretty != "":
		return d.Title.Pretty
	case d.Title.English != "":
		return d.Title.English
	default:
		return d.Title.Japanese
	}
}

// UnmarshalJSON implements [json.Unmarshaler].
func (d *Doujinshi) UnmarshalJSON(data []byte) error {
	type Alias Doujinshi
	correctStruct := struct {
		*Alias
		ID         json.Number   `json:"id"`
		MediaID    json.Number   `json:"media_id"`
		UploadDate JSONTimestamp `json:"upload_date"`
	}{}

	err := json.Unmarshal(data, &correctStruct)
	if err != nil {
		return err
	}

	ID, err := correctStruct.MediaID.Int64()
	if err != nil {
		return fmt.Errorf("failed to parse ID from '%s': %w", correctStruct.MediaID, err)
	}
	d.ID = int(ID)

	mediaID, err := correctStruct.MediaID.Int64()
	if err != nil {
		return fmt.Errorf("failed to parse media ID from '%s': %w", correctStruct.MediaID, err)
	}
	d.MediaID = int(mediaID)

	d.Title = correctStruct.Title
	d.Tags = correctStruct.Tags
	d.Scanlator = correctStruct.Scanlator
	d.NumPages = correctStruct.NumPages
	d.NumFavorites = correctStruct.NumFavorites
	d.UploadDate = correctStruct.UploadDate.Time
	d.Images = correctStruct.Images

	return nil
}
