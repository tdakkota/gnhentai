package testutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/nhentaiapi"
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

		h, err := c.Search(ctx, "ahegao", 1)
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

		h, err := c.SearchByTag(ctx, nhentaiapi.Tag{Name: "milf", ID: 1207}, 1)
		require.NoError(t, err)
		require.NotEmpty(t, h)
	})

	t.Run("SecondPage", func(t *testing.T) {
		c := constructor(t)

		h, err := c.SearchByTag(ctx, nhentaiapi.Tag{Name: "milf", ID: 1207}, 2)
		require.NoError(t, err)
		require.NotEmpty(t, h)
	})

	t.Run("TagDoesNotExist", func(t *testing.T) {
		c := constructor(t)

		_, err := c.SearchByTag(ctx, nhentaiapi.Tag{Name: "lolkek"}, 1)
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

		d, err := c.Related(ctx, "305329")
		require.NoError(t, err)
		require.NotEmpty(t, d.Result)

		t.Log(d.Result[0].Title.English)
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		c := constructor(t)
		_, err := c.Related(ctx, "-100500")
		require.Error(t, err, "id does not exist")
	})
}
