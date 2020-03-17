package gnhentai

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
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

	dataFile, err := os.Open("testdata/integration.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer dataFile.Close()

	var testData map[string]Doujinshi
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

			h, err := c.ByID(numberID)
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

			require.Equal(t, data, d)
		})
	}
}
