package parser

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/tdakkota/gnhentai"
)

// ParseComic parses a comic from a HTML document.
func ParseComic(sel *goquery.Selection) (gnhentai.Doujinshi, error) {
	return parseComic(sel)
}

func parseComic(doc *goquery.Selection) (result gnhentai.Doujinshi, err error) {
	infoBlock := doc.Find("#info")
	result.Title.English = infoBlock.Find("h1").First().Text()
	result.Title.Japanese = infoBlock.Find("h2").First().Text()

	uploaded := infoBlock.Find("div time").First()
	if datetime, ok := uploaded.Attr("datetime"); ok {
		result.UploadDate, err = time.Parse(time.RFC3339Nano, datetime)
		if err != nil {
			return result, fmt.Errorf("parse timestamp %q: %w", datetime, err)
		}
	}

	allTags := infoBlock.Find("#tags").First().Children()
	allTags.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var (
			tags    []gnhentai.Tag
			tagType string
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

	if link, ok := absoluteBaseLink(doc.Find("#cover a"), "href"); ok {
		u, err := url.Parse(link)
		if err != nil {
			return result, fmt.Errorf("parse cover link in %q: %w", link, err)
		}

		var n int
		_, err = fmt.Sscanf(u.Path, "/g/%d/%d/", &result.ID, &n)
		if err != nil {
			return result, fmt.Errorf("parse cover link in %q: %w", link, err)
		}
	}

	if link, ok := absoluteBaseLink(doc.Find("#cover a img"), "data-src"); ok {
		u, err := url.Parse(link)
		if err != nil {
			return result, fmt.Errorf("parse cover link in %q: %w", link, err)
		}

		_, err = fmt.Sscanf(u.Path, "/galleries/%d/cover", &result.MediaID)
		if err != nil {
			return result, fmt.Errorf("parse cover link in %q: %w", link, err)
		}
	}

	return result, nil
}

// ErrNoID is returned when there is no ID to parse.
var ErrNoID = errors.New("no ID to parse")

func absoluteBaseLink(a *goquery.Selection, attrName string) (string, bool) {
	if link, ok := a.Attr(attrName); ok {
		link = strings.TrimSpace(link)
		if strings.Index(link, "https://") != 0 {
			link = gnhentai.BaseNHentaiLink + link
		}
		return link, true
	}
	return "", false
}

func parseTag(link *goquery.Selection) (result gnhentai.Tag, err error) {
	countNode := link.Find(".count").First()
	counterText := strings.ReplaceAll(countNode.Text(), ",", "")
	counterText = strings.Trim(counterText, "()")

	_, err = fmt.Sscanf(counterText, "%d", &result.Count)
	if err != nil {
		return result, fmt.Errorf("parse counter %q: %w", counterText, err)
	}
	countNode.Remove()

	result.Name = strings.TrimSpace(link.Text())
	if tagLink, ok := absoluteBaseLink(link, "href"); ok {
		result.URL = tagLink
	}

	if class, ok := link.Attr("class"); ok {
		_, err = fmt.Sscanf(class, "tag tag-%d", &result.ID)
		if err != nil {
			return result, fmt.Errorf("parse ID %q: %w", class, err)
		}
	} else {
		return result, ErrNoID
	}

	return result, nil
}

func parseTags(t string, tags *goquery.Selection) (result []gnhentai.Tag, err error) {
	result = make([]gnhentai.Tag, 0, tags.Children().Length())
	tags.Children().EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var tag gnhentai.Tag

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

func mapTagType(name string) (string, bool) {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "parodies:":
		return "parody", true
	case "characters:":
		return "character", true
	case "tags:":
		return "tag", true
	case "artists:":
		return "artist", true
	case "groups:":
		return "group", true
	case "languages:":
		return "language", true
	case "categories:":
		return "category", true
	default:
		return "", false
	}
}

// ParseSearch parses search results from a HTML document.
func ParseSearch(sel *goquery.Selection) ([]gnhentai.Doujinshi, error) {
	return parseSearch(sel)
}

// ParseRelated parses related doujinshi from a HTML document.
func ParseRelated(sel *goquery.Selection) ([]gnhentai.Doujinshi, error) {
	return parseSearch(sel.Find("#related-container"))
}

func parseSearch(doc *goquery.Selection) (result []gnhentai.Doujinshi, err error) {
	galleries := doc.Find(".gallery")
	result = make([]gnhentai.Doujinshi, 0, galleries.Length())

	galleries.EachWithBreak(func(i int, gallery *goquery.Selection) bool {
		var doujinshi gnhentai.Doujinshi

		doujinshi, err = parseGallery(gallery)
		if err != nil {
			return false
		}

		result = append(result, doujinshi)
		return true
	})

	return result, nil
}

func parseGallery(gallery *goquery.Selection) (result gnhentai.Doujinshi, err error) {
	if link, ok := absoluteBaseLink(gallery.Find("a").First(), "href"); ok {
		_, err = fmt.Sscanf(link, "https://nhentai.net/g/%d/", &result.ID)
		if err != nil {
			return result, fmt.Errorf("parse doujinshi link %q: %w", link, err)
		}
	}

	if link, ok := absoluteBaseLink(gallery.Find("img").First(), "data-src"); ok {
		_, err = fmt.Sscanf(link, "https://t.nhentai.net/galleries/%d/thumb", &result.MediaID)
		if err != nil {
			return result, fmt.Errorf("parse cover link in %q: %w", link, err)
		}
	}

	result.Title.English = gallery.Find(".caption").First().Text()
	result.Title.Japanese = result.Title.English
	return
}
