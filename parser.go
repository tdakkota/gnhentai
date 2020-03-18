package gnhentai

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

func parse(doc *goquery.Document) (result Doujinshi, err error) {
	infoBlock := doc.Find("#info")
	result.Name = infoBlock.Find("h1").First().Text()
	result.AlterName = infoBlock.Find("h2").First().Text()

	uploaded := infoBlock.Find("div time").First()
	if datetime, ok := uploaded.Attr("datetime"); ok {
		result.Uploaded, err = time.Parse(time.RFC3339Nano, datetime)
		if err != nil {
			return result, fmt.Errorf("failed to parse timestamp: %v", err)
		}
	}

	allTags := infoBlock.Find("#tags").First().Children()
	allTags.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var tags []BaseTag

		tags, err = parseBaseTags(selection.Find(".tags"))
		if err != nil {
			return false
		}

		if len(selection.Nodes) > 0 {
			textNode := selection.Nodes[0].FirstChild
			if textNode != nil {
				err = mapTagType(&result, textNode.Data, tags)
				if err != nil {
					return false
				}
			}
		}

		return true
	})

	result.Previews, result.MediaID, err = parsePreviews(doc.Find("#thumbnail-container").First())
	if err != nil {
		return result, err
	}

	return result, nil
}

var ErrNoID = errors.New("no ID to parse")

func absoluteBaseLink(a *goquery.Selection) (string, bool) {
	if link, ok := a.Attr("href"); ok {
		link = strings.TrimSpace(link)
		if strings.Index(link, BaseNHentaiLink) != 0 {
			link = BaseNHentaiLink + link
		}
		return link, true
	}
	return "", false
}

func parseBaseTag(link *goquery.Selection) (result BaseTag, err error) {
	countNode := link.Find(".count").First()

	counterText := strings.Join(strings.Split(countNode.Text(), ","), "")
	_, err = fmt.Sscanf(counterText, "(%d)", &result.Count)
	if err != nil {
		return result, fmt.Errorf("failed to parse counter: %v", err)
	}
	countNode.Remove()

	result.Name = strings.TrimSpace(link.Text())
	if tagLink, ok := absoluteBaseLink(link); ok {
		result.Link = tagLink
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

func parseBaseTags(tags *goquery.Selection) ([]BaseTag, error) {
	var err error
	result := make([]BaseTag, 0, tags.Children().Length())

	tags.Children().EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var tag BaseTag

		tag, err = parseBaseTag(selection)
		if err != nil {
			return false
		}
		result = append(result, tag)

		return true
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func mapTagType(result *Doujinshi, name string, tags []BaseTag) error {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "parodies:":
		result.Parodies = tags
	case "characters:":
		result.Characters = tags
	case "tags:":
		result.Tags = tags
	case "artists:":
		result.Artists = tags
	case "groups:":
		result.Groups = tags
	case "languages:":
		result.Languages = tags
	case "categories:":
		result.Categories = tags
	default:
		return fmt.Errorf("unknown tag type: %s", name)
	}

	return nil
}

func parsePreviews(container *goquery.Selection) (Previews, int, error) {
	thumbs := container.Find(".gallerythumb")

	result := make(Previews, 0, thumbs.Length())
	mediaID := 0
	var err error

	thumbs.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		preview := Preview{Number: i + 1}

		if link, ok := absoluteBaseLink(selection); ok {
			preview.Link = link
			if mediaID == 0 {
				_, err = fmt.Sscanf(link, "https://nhentai.net/g/%d/%d/", &mediaID, &preview.Number)
				if err != nil {
					err = fmt.Errorf("failed to parse thumb %d link in '%s': %v", i+1, link, err)
					return false
				}
			}
		}

		img := selection.Find("img").First()

		if link, ok := img.Attr("data-src"); ok {
			preview.Thumbnail.Link = link
		}

		if preview.Thumbnail.Width, err = strconv.Atoi(img.AttrOr("width", "0")); err != nil {
			return false
		}

		if preview.Thumbnail.Height, err = strconv.Atoi(img.AttrOr("height", "0")); err != nil {
			return false
		}

		result = append(result, preview)
		return true
	})

	return result, mediaID, err
}
