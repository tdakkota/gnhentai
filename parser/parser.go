package parser

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tdakkota/gnhentai"
	"io"
	"strings"
	"time"
)

func Parse(r io.Reader) (result gnhentai.Doujinshi, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return
	}
	return parse(doc)
}

func parse(doc *goquery.Document) (result gnhentai.Doujinshi, err error) {
	infoBlock := doc.Find("#info")
	result.Title.English = infoBlock.Find("h1").First().Text()
	result.Title.Japanese = infoBlock.Find("h2").First().Text()

	uploaded := infoBlock.Find("div time").First()
	if datetime, ok := uploaded.Attr("datetime"); ok {
		result.UploadDate, err = time.Parse(time.RFC3339Nano, datetime)
		if err != nil {
			return result, fmt.Errorf("failed to parse timestamp: %v", err)
		}
	}

	allTags := infoBlock.Find("#tags").First().Children()
	allTags.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var tags []gnhentai.Tag
		var tagType string

		if len(selection.Nodes) > 0 {
			textNode := selection.Nodes[0].FirstChild
			if textNode != nil {
				tagType, err = mapTagType(textNode.Data)
				if err != nil {
					return false
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

	if link, ok := absoluteBaseLink(doc.Find("#cover a")); ok {
		var n int
		_, err = fmt.Sscanf(link, "https://nhentai.net/g/%d/%d/", &result.MediaID, &n)
		if err != nil {
			return result, fmt.Errorf("failed to parse cover link in '%s': %v", link, err)
		}
	}

	return result, nil
}

var ErrNoID = errors.New("no ID to parse")

func absoluteBaseLink(a *goquery.Selection) (string, bool) {
	if link, ok := a.Attr("href"); ok {
		link = strings.TrimSpace(link)
		if strings.Index(link, gnhentai.BaseNHentaiLink) != 0 {
			link = gnhentai.BaseNHentaiLink + link
		}
		return link, true
	}
	return "", false
}

func parseTag(link *goquery.Selection) (result gnhentai.Tag, err error) {
	countNode := link.Find(".count").First()

	counterText := strings.Join(strings.Split(countNode.Text(), ","), "")
	_, err = fmt.Sscanf(counterText, "(%d)", &result.Count)
	if err != nil {
		return result, fmt.Errorf("failed to parse counter: %v", err)
	}
	countNode.Remove()

	result.Name = strings.TrimSpace(link.Text())
	if tagLink, ok := absoluteBaseLink(link); ok {
		result.URL = tagLink
	}

	if class, ok := link.Attr("class"); ok {
		_, err = fmt.Sscanf(class, "tag tag-%d", &result.ID)
		if err != nil {
			return result, fmt.Errorf("failed to parse ID: %v", err)
		}
	} else {
		return result, ErrNoID
	}

	return result, nil
}

func parseTags(t string, tags *goquery.Selection) ([]gnhentai.Tag, error) {
	var err error
	result := make([]gnhentai.Tag, 0, tags.Children().Length())

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

	if err != nil {
		return nil, err
	}
	return result, nil
}

func mapTagType(name string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "parodies:":
		return "parody", nil
	case "characters:":
		return "character", nil
	case "tags:":
		return "tag", nil
	case "artists:":
		return "artist", nil
	case "groups:":
		return "group", nil
	case "languages:":
		return "language", nil
	case "categories:":
		return "category", nil
	default:
		return "", fmt.Errorf("unknown tag type: %s", name)
	}
}
