package parser

import (
	"bytes"
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
	"github.com/tdakkota/gnhentai"
)

func Test_parseBaseTag(t *testing.T) {
	data := map[string]gnhentai.Tag{
		`<a href="/group/applique/" class="tag tag-109391 ">applique <span class="count">(2)</span></a>`: {
			ID:    109391,
			Count: 2,
			Name:  "applique",
			URL:   "https://nhentai.net/group/applique/",
		},
	}

	i := 0
	for html, expected := range data {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(html)
			doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(html))
			if err != nil {
				t.Error(err)
				return
			}

			tag, err := parseTag(doc.Find("a"))
			if err != nil {
				t.Error(err)
				return
			}

			t.Log(tag)
			require.Equal(t, expected, tag)
		})
		i++
	}
}

func Test_parseTags(t *testing.T) {
	data := map[string][]gnhentai.Tag{
		`<span class="tags">
			<a href="/tag/big-breasts/" class="tag tag-2937 ">big breasts <span class="count">(101,029)</span></a>
			<a href="/tag/full-color/" class="tag tag-20905 ">full color <span class="count">(31,918)</span></a>
			<a href="/tag/full-censorship/" class="tag tag-8368 ">full censorship <span class="count">(16,230)</span></a>
			<a href="/tag/big-ass/" class="tag tag-9083 ">big ass <span class="count">(9,317)</span></a>
			<a href="/tag/webtoon/" class="tag tag-50585 ">webtoon <span class="count">(1,618)</span></a>
		</span>`: {
			{
				ID:    2937,
				Count: 101029,
				Name:  "big breasts",
				URL:   "https://nhentai.net/tag/big-breasts/",
			},
			{
				ID:    20905,
				Count: 31918,
				Name:  "full color",
				URL:   "https://nhentai.net/tag/full-color/",
			},
			{
				ID:    8368,
				Count: 16230,
				Name:  "full censorship",
				URL:   "https://nhentai.net/tag/full-censorship/",
			},
			{
				ID:    9083,
				Count: 9317,
				Name:  "big ass",
				URL:   "https://nhentai.net/tag/big-ass/",
			},
			{
				ID:    50585,
				Count: 1618,
				Name:  "webtoon",
				URL:   "https://nhentai.net/tag/webtoon/",
			},
		},
	}

	i := 0
	for html, expect := range data {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(html)
			doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(html))
			if err != nil {
				t.Error(err)
				return
			}

			tags, err := parseTags("", doc.Find(".tags"))
			if err != nil {
				t.Error(err)
				return
			}

			require.NotEmpty(t, tags)
			for _, tag := range tags {
				t.Log(tag.ID, tag.Name, tag.Count)
			}

			require.Equal(t, expect, tags)
		})
		i++
	}
}

func TestParse(t *testing.T) {
	files := map[string]gnhentai.Doujinshi{
		"../testdata/test.html":  {},
		"../testdata/test2.html": {},
	}

	for fileName := range files {
		t.Run(fileName, func(t *testing.T) {
			f, err := os.Open(fileName)
			if err != nil {
				t.Error(err)
				return
			}

			doc, err := goquery.NewDocumentFromReader(f)
			if err != nil {
				return
			}

			r, err := parseComic(doc.Selection)
			if err != nil {
				t.Error(err)
				return
			}

			d, err := json.MarshalIndent(r, "", "\t")
			if err != nil {
				t.Error(err)
				return
			}

			t.Log(string(d))
		})
	}
}

func TestSearchParse(t *testing.T) {
	files := map[string][]gnhentai.Doujinshi{
		"../testdata/test3.html": nil,
	}

	for fileName := range files {
		t.Run(fileName, func(t *testing.T) {
			f, err := os.Open(fileName)
			if err != nil {
				t.Error(err)
				return
			}

			doc, err := goquery.NewDocumentFromReader(f)
			if err != nil {
				return
			}

			r, err := parseSearch(doc.Selection)
			if err != nil {
				t.Error(err)
				return
			}

			d, err := json.MarshalIndent(r, "", "\t")
			if err != nil {
				t.Error(err)
				return
			}

			t.Log(string(d))
		})
	}
}
