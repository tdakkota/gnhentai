package testutil

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tdakkota/gnhentai"
)

// Skip test if E2E env is not set.
func Skip(t testing.TB) {
	t.Helper()
	if os.Getenv("E2E") == "" {
		t.Skip("Set E2E env to run")
	}
}

// TestRandom tests [gnhentai.Client.Random].
func TestRandom(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	Skip(t)
	t.Helper()
	ctx := t.Context()

	c := constructor(t)
	h, err := c.Random(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, h.Name())
}

// TestSearch tests [gnhentai.Client.Search].
func TestSearch(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	Skip(t)
	t.Helper()
	ctx := t.Context()

	t.Run("FirstPage", func(t *testing.T) {
		c := constructor(t)

		h, err := c.Search(ctx, "ahegao", 0)
		require.NoError(t, err)
		require.NotEmpty(t, h)
	})

	t.Run("SecondPage", func(t *testing.T) {
		c := constructor(t)

		h, err := c.Search(ctx, "ahegao", 2)
		require.NoError(t, err)
		require.NotEmpty(t, h)
	})
}

// TestSearchByTag tests [gnhentai.Client.SearchByTag].
func TestSearchByTag(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	Skip(t)
	t.Helper()
	ctx := t.Context()

	t.Run("FirstPage", func(t *testing.T) {
		c := constructor(t)

		h, err := c.SearchByTag(ctx, gnhentai.Tag{Name: "milf", ID: 1207}, 0)
		require.NoError(t, err)
		require.NotEmpty(t, h)
	})

	t.Run("SecondPage", func(t *testing.T) {
		c := constructor(t)

		h, err := c.SearchByTag(ctx, gnhentai.Tag{Name: "milf", ID: 1207}, 2)
		require.NoError(t, err)
		require.NotEmpty(t, h)
	})

	t.Run("TagDoesNotExist", func(t *testing.T) {
		c := constructor(t)

		_, err := c.SearchByTag(ctx, gnhentai.Tag{Name: "lolkek"}, 0)
		require.Error(t, err, "tag does not exist")
	})
}

// TestRelated tests [gnhentai.Client.Related].
func TestRelated(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	Skip(t)
	t.Helper()
	ctx := t.Context()

	t.Run("Get", func(t *testing.T) {
		c := constructor(t)

		d, err := c.Related(ctx, 305329)
		require.NoError(t, err)
		require.NotEmpty(t, d)

		t.Log(d[0].Title.English)
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		c := constructor(t)
		_, err := c.Related(ctx, -100500)
		require.Error(t, err, "id does not exist")
	})
}

// TestGetByID tests [gnhentai.Client.GetByID].
func TestGetByID(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	Skip(t)
	t.Helper()
	ctx := t.Context()

	dataFile, err := os.Open("../testdata/integration.json")
	require.NoError(t, err)
	defer func() {
		_ = dataFile.Close()
	}()

	var testData map[string]gnhentai.Doujinshi
	require.NoError(t, json.NewDecoder(dataFile).Decode(&testData))

	for id, data := range testData {
		numberID, err := strconv.Atoi(id)
		require.NoError(t, err)

		t.Run(id, func(t *testing.T) {
			c := constructor(t)

			doujinshi, err := c.ByID(ctx, numberID)
			require.NoError(t, err)

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

	t.Run("DoesNotExists", func(t *testing.T) {
		c := constructor(t)

		_, err := c.ByID(ctx, -100500)
		require.Error(t, err, "id does not exist")
	})
}

// RunAll runs all tests.
func RunAll(t *testing.T, constructor func(t *testing.T) gnhentai.Client) {
	for _, tt := range []struct {
		name string
		f    func(t *testing.T, constructor func(t *testing.T) gnhentai.Client)
	}{
		{"TestRandom", TestRandom},
		{"TestSearch", TestSearch},
		{"TestSearchByTag", TestSearchByTag},
		{"TestRelated", TestRelated},
		{"TestGetByID", TestGetByID},
	} {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			tc.f(t, constructor)
		})
	}
}
