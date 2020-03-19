package testutil

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/tdakkota/gnhentai"
	"os"
	"strconv"
	"testing"
)

func TestRandom(t *testing.T, constructor func() gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	c := constructor()
	h, err := c.Random()
	if err != nil {
		t.Error(err)
		return
	}

	d, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(d))
}

func TestSearch(t *testing.T, constructor func() gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("first-page", func(t *testing.T) {
		c := constructor()
		h, err := c.Search("ahegao", 0)
		if err != nil {
			t.Error(err)
			return
		}

		d, err := json.MarshalIndent(h, "", "\t")
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(string(d))
	})

	t.Run("second-page", func(t *testing.T) {
		c := constructor()
		h, err := c.Search("ahegao", 2)
		if err != nil {
			t.Error(err)
			return
		}

		d, err := json.MarshalIndent(h, "", "\t")
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(string(d))
	})
}

func TestSearchByTag(t *testing.T, constructor func() gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("first-page", func(t *testing.T) {
		c := constructor()
		h, err := c.SearchByTag(gnhentai.Tag{Name: "milf"}, 0)
		if err != nil {
			t.Error(err)
			return
		}

		d, err := json.MarshalIndent(h, "", "\t")
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(string(d))
	})

	t.Run("second-page", func(t *testing.T) {
		c := constructor()
		h, err := c.SearchByTag(gnhentai.Tag{Name: "milf"}, 2)
		if err != nil {
			t.Error(err)
			return
		}

		d, err := json.MarshalIndent(h, "", "\t")
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(string(d))
	})

	t.Run("tag-does-not-exists", func(t *testing.T) {
		c := constructor()
		_, err := c.SearchByTag(gnhentai.Tag{Name: "lolkek"}, 0)
		if err == nil {
			t.Error("tag does not exists - should return error")
			return
		} else {
			t.Log(err)
		}
	})
}

func TestRelated(t *testing.T, constructor func() gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("get from 305329", func(t *testing.T) {
		c := constructor()

		d, err := c.Related(305329)
		if err != nil {
			t.Error(err)
			return
		}

		require.NotEmpty(t, d)
		t.Log(d[0].Title.English)
	})

	t.Run("id-does-not-exists", func(t *testing.T) {
		c := constructor()
		_, err := c.Related(-100500)
		if err == nil {
			t.Error("id does not exists - should return error")
			return
		} else {
			t.Log(err)
		}
	})
}

func TestGetByID(t *testing.T, constructor func() gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

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
			c := constructor()

			doujinshi, err := c.ByID(numberID)
			if err != nil {
				t.Error(err)
				return
			}

			pretty, err := json.MarshalIndent(doujinshi, "", "\t")
			if err != nil {
				t.Error(err)
				return
			}

			t.Log(string(pretty))

			require.Equal(t, data.Title.Pretty, doujinshi.Title.Pretty)
			require.Equal(t, data.Title.English, doujinshi.Title.English)
			require.Equal(t, data.Title.Japanese, doujinshi.Title.Japanese)
			require.NotEmpty(t, data.Tags)
			require.Equal(t, data.Tags[0].Name, doujinshi.Tags[0].Name)
			require.Equal(t, data.NumPages, doujinshi.NumPages)
		})
	}
}

func TestGetByID2(t *testing.T, constructor func() gnhentai.Client) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("id-does-not-exists", func(t *testing.T) {
		c := constructor()
		_, err := c.ByID(-100500)
		if err == nil {
			t.Error("id does not exists - should return error")
			return
		} else {
			t.Log(err)
		}
	})
}

func RunAll(t *testing.T, constructor func() gnhentai.Client) {
	funcs := map[string]func(t *testing.T, constructor func() gnhentai.Client){
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
