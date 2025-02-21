package parser

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-faster/sdk/gold"
	"github.com/stretchr/testify/require"

	"github.com/tdakkota/gnhentai"
)

func Test_parseBaseTag(t *testing.T) {
	for i, tt := range []struct {
		html string
		want gnhentai.Tag
	}{
		{
			`<a href="/group/applique/" class="tag tag-109391 ">applique <span class="count">(2)</span></a>`,
			gnhentai.Tag{
				ID:    109391,
				Count: 2,
				Name:  "applique",
				URL:   "https://nhentai.net/group/applique/",
			},
		},
	} {
		tt := tt

		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			require.NoError(t, err)

			tag, err := parseTag(doc.Find("a"))
			require.NoError(t, err)
			require.Equal(t, tt.want, tag)
		})
	}
}

func Test_parseTags(t *testing.T) {
	for i, tt := range []struct {
		html string
		want []gnhentai.Tag
	}{
		{
			`<span class="tags">
			<a href="/tag/big-breasts/" class="tag tag-2937 ">big breasts <span class="count">(101,029)</span></a>
			<a href="/tag/full-color/" class="tag tag-20905 ">full color <span class="count">(31,918)</span></a>
			<a href="/tag/full-censorship/" class="tag tag-8368 ">full censorship <span class="count">(16,230)</span></a>
			<a href="/tag/big-ass/" class="tag tag-9083 ">big ass <span class="count">(9,317)</span></a>
			<a href="/tag/webtoon/" class="tag tag-50585 ">webtoon <span class="count">(1,618)</span></a>
		</span>`,
			[]gnhentai.Tag{
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
		},
	} {
		tt := tt

		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			require.NoError(t, err)

			tags, err := parseTags("", doc.Find(".tags"))
			require.NoError(t, err)
			require.NotEmpty(t, tags)
			require.Equal(t, tt.want, tags)
		})
	}
}

//go:embed _testdata
var testdata embed.FS

func TestParse(t *testing.T) {
	const testType = "comic"
	testdataDir := path.Join("_testdata", testType)
	files, err := fs.ReadDir(testdata, testdataDir)
	require.NoError(t, err)

	for _, fi := range files {
		name := fi.Name()
		t.Run(name, func(t *testing.T) {
			data, err := testdata.Open(path.Join(testdataDir, name))
			require.NoError(t, err)
			defer data.Close()

			doc, err := goquery.NewDocumentFromReader(data)
			require.NoError(t, err)

			r, err := parseComic(doc.Selection)
			require.NoError(t, err)

			d, err := json.MarshalIndent(r, "", "\t")
			require.NoError(t, err)

			gold.Str(t, string(d), testType, name+".json")
		})
	}
}

func TestSearchParse(t *testing.T) {
	const testType = "search"
	testdataDir := path.Join("_testdata", testType)
	files, err := fs.ReadDir(testdata, testdataDir)
	require.NoError(t, err)

	for _, fi := range files {
		name := fi.Name()
		t.Run(name, func(t *testing.T) {
			data, err := testdata.Open(path.Join(testdataDir, name))
			require.NoError(t, err)
			defer data.Close()

			doc, err := goquery.NewDocumentFromReader(data)
			require.NoError(t, err)

			r, err := parseSearch(doc.Selection)
			require.NoError(t, err)

			d, err := json.MarshalIndent(r, "", "\t")
			require.NoError(t, err)

			gold.Str(t, string(d), testType, name+".json")
		})
	}
}
