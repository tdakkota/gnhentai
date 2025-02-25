package parser

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/dustin/go-humanize"
	"github.com/go-faster/jx"

	"github.com/tdakkota/gnhentai/nhentaiapi"
)

// ParseComic parses a comic from a HTML document.
func ParseComic(sel *goquery.Selection) (*nhentaiapi.Book, error) {
	if r, ok := parseScriptJSON[nhentaiapi.Book](sel); ok {
		for i := range r.Tags {
			tag := &r.Tags[i]

			u, ok := tag.URL.Get()
			if !ok {
				continue
			}

			ref, err := url.Parse(u)
			if err != nil {
				continue
			}

			tag.URL.SetTo(nhentaiapi.BaseNHentaiLink.ResolveReference(ref).String())
		}
		return &r, nil
	}
	return parseComic(sel)
}

func parseComic(doc *goquery.Selection) (result *nhentaiapi.Book, err error) {
	result = new(nhentaiapi.Book)

	if id, ok := strings.CutPrefix(doc.Find("#gallery_id").Text(), "#"); ok {
		parseBookID(id, &result.ID)
	}

	if link, ok := absoluteBaseLink(doc.Find("#cover a img"), "data-src"); ok {
		u, err := url.Parse(link)
		if err != nil {
			return result, fmt.Errorf("parse cover URL %q: %w", link, err)
		}
		// Cut off the filename and get media ID.
		result.MediaID = path.Base(path.Dir(u.Path))
	}

	infoBlock := doc.Find("#info")
	if text := infoBlock.Find("h1").First().Text(); text != "" {
		result.Title.English.SetTo(text)
	}
	if text := infoBlock.Find("h2").First().Text(); text != "" {
		result.Title.Japanese.SetTo(text)
	}

	uploaded := infoBlock.Find("div time").First()
	if datetime, ok := uploaded.Attr("datetime"); ok {
		t, err := time.Parse(time.RFC3339Nano, datetime)
		if err != nil {
			return result, fmt.Errorf("parse timestamp %q: %w", datetime, err)
		}
		result.UploadDate.SetTo(t)
	}

	allTags := infoBlock.Find("#tags").First().Children()
	allTags.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var (
			tags    []nhentaiapi.Tag
			tagType nhentaiapi.TagType
		)

		if len(selection.Nodes) > 0 {
			textNode := selection.Nodes[0].FirstChild
			if textNode != nil {
				// Gracefully skip unknown tag type.
				var ok bool
				tagType, ok = mapTagType(textNode.Data)
				if !ok {
					return true
				}
			}
		}

		tags, err = parseTags(tagType, selection.Find(".tags"))
		if err != nil {
			return false
		}

		result.Tags = append(result.Tags, tags...)
		return true
	})
	if err != nil {
		return result, fmt.Errorf("parse tags: %w", err)
	}

	allThumbs := doc.Find("#thumbnail-container").Find(".gallerythumb")
	allThumbs.Each(func(i int, s *goquery.Selection) {
		src, ok := s.Find("img").Attr("data-src")
		if !ok {
			return
		}

		u, err := url.Parse(src)
		if err != nil {
			return
		}

		ext := path.Ext(path.Base(u.Path))
		if ext == "" {
			return
		}
		ext = strings.TrimPrefix(ext, ".")
		ext = strings.ToLower(ext)

		result.Images.Pages = append(result.Images.Pages, nhentaiapi.Image{
			T: ext[:1],
		})
	})
	result.NumPages.SetTo(len(result.Images.Pages))

	return result, nil
}

func parseTags(t nhentaiapi.TagType, tags *goquery.Selection) (result []nhentaiapi.Tag, err error) {
	result = make([]nhentaiapi.Tag, 0, tags.Children().Length())
	tags.Children().EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var tag nhentaiapi.Tag

		tag, err = parseTag(selection)
		if err != nil {
			return false
		}
		tag.Type = t

		result = append(result, tag)

		return true
	})
	return result, err
}

// ErrNoID is returned when there is no ID to parse.
var ErrNoID = errors.New("no ID to parse")

func parseTag(link *goquery.Selection) (result nhentaiapi.Tag, err error) {
	countNode := link.Find(".count").First()
	counterText := strings.ReplaceAll(countNode.Text(), ",", "")
	counterText = strings.Trim(counterText, "()")

	counter, _, err := humanize.ParseSI(counterText)
	if err == nil {
		result.Count.SetTo(int(counter))
	}
	countNode.Remove()

	result.Name = strings.TrimSpace(link.Text())
	if tagLink, ok := absoluteBaseLink(link, "href"); ok {
		result.URL.SetTo(tagLink)
	}

	if class, ok := link.Attr("class"); ok {
		_, err = fmt.Sscanf(class, "tag tag-%d", &result.ID)
		if err != nil {
			return result, fmt.Errorf("parse tag ID %q: %w", class, err)
		}
	} else {
		return result, ErrNoID
	}

	return result, nil
}

func mapTagType(name string) (nhentaiapi.TagType, bool) {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "parodies:":
		return nhentaiapi.TagTypeParody, true
	case "characters:":
		return nhentaiapi.TagTypeCharacter, true
	case "tags:":
		return nhentaiapi.TagTypeTag, true
	case "artists:":
		return nhentaiapi.TagTypeArtist, true
	case "groups:":
		return nhentaiapi.TagTypeGroup, true
	case "languages:":
		return nhentaiapi.TagTypeCategory, true
	case "categories:":
		return nhentaiapi.TagTypeLanguage, true
	default:
		return "", false
	}
}

// ParseSearch parses search results from a HTML document.
func ParseSearch(sel *goquery.Selection) (*nhentaiapi.SearchResponse, error) {
	return parseSearch(sel)
}

// ParseRelated parses related doujinshi from a HTML document.
func ParseRelated(sel *goquery.Selection) (*nhentaiapi.SearchResponse, error) {
	return parseSearch(sel.Find("#related-container"))
}

func parseSearch(doc *goquery.Selection) (result *nhentaiapi.SearchResponse, err error) {
	galleries := doc.Find(".gallery")
	result = &nhentaiapi.SearchResponse{
		Result: make([]nhentaiapi.Book, 0, galleries.Length()),
	}

	galleries.EachWithBreak(func(i int, gallery *goquery.Selection) bool {
		var doujinshi nhentaiapi.Book

		doujinshi, err = parseGallery(gallery)
		if err != nil {
			return false
		}

		result.Result = append(result.Result, doujinshi)
		return true
	})

	return result, nil
}

func parseGallery(gallery *goquery.Selection) (result nhentaiapi.Book, err error) {
	if link, ok := absoluteBaseLink(gallery.Find("a").First(), "href"); ok {
		u, err := url.Parse(link)
		if err != nil {
			return result, fmt.Errorf("parse doujinshi URL %q: %w", link, err)
		}
		parseBookID(path.Base(u.Path), &result.ID)
	}

	if link, ok := absoluteBaseLink(gallery.Find("img").First(), "data-src"); ok {
		u, err := url.Parse(link)
		if err != nil {
			return result, fmt.Errorf("parse cover URL %q: %w", link, err)
		}
		// Cut off the filename and get media ID.
		result.MediaID = path.Base(path.Dir(u.Path))
	}

	if text := gallery.Find(".caption").First().Text(); text != "" {
		result.Title.English.SetTo(text)
		result.Title.Japanese.SetTo(text)
	}
	return
}

func parseBookID(value string, bookID *nhentaiapi.BookID) {
	if v, err := strconv.Atoi(value); err == nil {
		bookID.SetInt(v)
	} else {
		bookID.SetString(value)
	}
}

func absoluteBaseLink(a *goquery.Selection, attrName string) (string, bool) {
	link, ok := a.Attr(attrName)
	if !ok {
		return "", false
	}
	link = strings.TrimSpace(link)
	return absoluteURL(link)
}

func absoluteURL(raw string) (string, bool) {
	ref, err := url.Parse(raw)
	if err != nil {
		return "", false
	}
	return nhentaiapi.BaseNHentaiLink.ResolveReference(ref).String(), true
}

var scriptRe = regexp.MustCompile(`window._gallery.+JSON.parse\((.*)\)`)

func parseScriptJSON[
	T any,
	P interface {
		*T
		Decode(*jx.Decoder) error
		Encode(*jx.Encoder)
	},
](doc *goquery.Selection) (r T, found bool) {
	doc.Find("script").EachWithBreak(func(i int, s *goquery.Selection) bool {
		matches := scriptRe.FindStringSubmatch(s.Text())
		if len(matches) < 2 {
			return true
		}

		unquoted, err := jx.DecodeStr(matches[1]).StrAppend(nil)
		if err != nil {
			return true
		}
		var zero T
		if err := P(&zero).Decode(jx.DecodeBytes(unquoted)); err != nil {
			fmt.Println(err)
			return true
		}
		r = zero
		found = true
		return false
	})
	return r, found
}
