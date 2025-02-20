package testutil

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tdakkota/gnhentai"
)

func TestRandom(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	c := constructor(t)
	h, err := c.Random()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(h.Title.English)
}

func TestSearch(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	t.Run("first-page", func(t *testing.T) {
		c := constructor(t)
		h, err := c.Search("ahegao", 0)
		if err != nil {
			t.Error(err)
			return
		}

		require.NotEmpty(t, h)
		t.Log(h[0].Title.English)
	})

	t.Run("second-page", func(t *testing.T) {
		c := constructor(t)
		h, err := c.Search("ahegao", 2)
		if err != nil {
			t.Error(err)
			return
		}

		require.NotEmpty(t, h)
		t.Log(h[0].Title.English)
	})
}

func TestSearchByTag(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	t.Run("first-page", func(t *testing.T) {
		c := constructor(t)
		h, err := c.SearchByTag(gnhentai.Tag{Name: "milf", ID: 1207}, 0)
		if err != nil {
			t.Error(err)
			return
		}

		require.NotEmpty(t, h)
		t.Log(h[0].Title.English)
	})

	t.Run("second-page", func(t *testing.T) {
		c := constructor(t)
		h, err := c.SearchByTag(gnhentai.Tag{Name: "milf", ID: 1207}, 2)
		if err != nil {
			t.Error(err)
			return
		}

		require.NotEmpty(t, h)
		t.Log(h[0].Title.English)
	})

	t.Run("tag-does-not-exists", func(t *testing.T) {
		c := constructor(t)
		_, err := c.SearchByTag(gnhentai.Tag{Name: "lolkek"}, 0)
		if err == nil {
			t.Error("tag does not exists - should return error")
			return
		}
		t.Log(err)
	})
}

func TestRelated(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	t.Run("get from 305329", func(t *testing.T) {
		c := constructor(t)

		d, err := c.Related(305329)
		if err != nil {
			t.Error(err)
			return
		}

		require.NotEmpty(t, d)
		t.Log(d[0].Title.English)
	})

	t.Run("id-does-not-exists", func(t *testing.T) {
		c := constructor(t)
		_, err := c.Related(-100500)
		if err == nil {
			t.Error("id does not exists - should return error")
			return
		}
		t.Log(err)
	})
}

func TestGetByID(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	dataFile, err := os.Open("../testdata/integration.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer dataFile.Close()

	var testData map[string]gnhentai.Doujinshi
	err = json.NewDecoder(dataFile).Decode(&testData)
	if err != nil {
		t.Error(err)
		return
	}

	for id, data := range testData {
		numberID, err := strconv.Atoi(id)
		if err != nil {
			t.Log("invalid testdata")
			t.Fail()
			return
		}

		t.Run(id, func(t *testing.T) {
			c := constructor(t)

			doujinshi, err := c.ByID(numberID)
			if err != nil {
				t.Error(err)
				return
			}

			t.Log("Title:", doujinshi.Title.English)

			d, err := json.MarshalIndent(doujinshi, "", "\t")
			if err != nil {
				t.Error(err)
				return
			}

			t.Log(string(d))

			require.Equal(t, data.Title.English, doujinshi.Title.English)
			require.NotEmpty(t, data.Tags)
			require.NotEmpty(t, doujinshi.Tags)

			names := make([]string, len(data.Tags))
			for i := range names {
				names[i] = data.Tags[i].Name
			}
			require.Contains(t, names, doujinshi.Tags[0].Name)
		})
	}
}

func TestGetByID2(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	t.Run("id-does-not-exists", func(t *testing.T) {
		c := constructor(t)
		_, err := c.ByID(-100500)
		if err == nil {
			t.Error("id does not exists - should return error")
			return
		}
		t.Log(err)
	})
}

func TestDownloader(t *testing.T, constructor func(t *testing.T) gnhentai.Downloader) {
	if testing.Short() {
		t.Skip("skipping integration test")
		return
	}
	t.Helper()

	t.Run("download-305329", func(t *testing.T) {
		d := constructor(t)
		err := gnhentai.DownloadAll(d, gnhentai.Doujinshi{MediaID: 1590084}, func(i int, d gnhentai.Doujinshi) string {
			return strconv.Itoa(i) + ".jpg"
		})

		if err != nil {
			t.Error(err)
			return
		}
	})
}

func RunAll(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	funcs := map[string]func(t *testing.T, constructor func(t *testing.T) gnhentai.Client){
		"TestRandom":      TestRandom,
		"TestSearch":      TestSearch,
		"TestSearchByTag": TestSearchByTag,
		"TestRelated":     TestRelated,
		"TestGetByID":     TestGetByID,
		"TestGetByID2":    TestGetByID2,
	}

	for name, f := range funcs {
		t.Run(name, func(t *testing.T) {
			f(t, constructor)
		})
	}
}
