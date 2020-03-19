package parser

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/tdakkota/gnhentai"
	"os"
	"strconv"
	"testing"
)

func TestRandom(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	c := NewClient()
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

func TestSearch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("first-page", func(t *testing.T) {
		c := NewClient()
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
		c := NewClient()
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

func TestSearchByTag(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("first-page", func(t *testing.T) {
		c := NewClient()
		h, err := c.SearchByTag("milf", 0)
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
		c := NewClient()
		h, err := c.SearchByTag("milf", 2)
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
		c := NewClient()
		_, err := c.SearchByTag("lolkek", 0)
		if err == nil {
			t.Error("tag does not exists - should return error")
			return
		} else {
			t.Log(err)
		}
	})
}

func TestGetByID(t *testing.T) {
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
			c := NewClient()

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
			require.Equal(t, data.Tags[0].Name, doujinshi.Tags[0].Name)
			require.Equal(t, data.NumPages, doujinshi.NumPages)
		})
	}
}

func TestGetByID2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("id-does-not-exists", func(t *testing.T) {
		c := NewClient()
		_, err := c.ByID(-100500)
		if err == nil {
			t.Error("id does not exists - should return error")
			return
		} else {
			t.Log(err)
		}
	})
}
