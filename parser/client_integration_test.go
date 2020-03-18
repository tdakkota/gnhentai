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
