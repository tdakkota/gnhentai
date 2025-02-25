package parser

import (
	"testing"

	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/internal/testutil"
)

func newParser(t *testing.T) gnhentai.Client {
	return NewParser(ParserOptions{
		Client: testutil.TestClient(t),
	})
}

func TestRandom(t *testing.T) {
	testutil.TestRandom(t, newParser)
}

func TestSearch(t *testing.T) {
	testutil.TestSearch(t, newParser)
}

func TestSearchByTag(t *testing.T) {
	testutil.TestSearchByTag(t, newParser)
}

func TestRelated(t *testing.T) {
	testutil.TestRelated(t, newParser)
}
