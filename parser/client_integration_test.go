package parser

import (
	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/testutil"
	"testing"
)

func newParser() gnhentai.Client {
	return NewParser()
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

func TestGetByID(t *testing.T) {
	testutil.TestGetByID(t, newParser)
}

func TestGetByID2(t *testing.T) {
	testutil.TestGetByID2(t, newParser)
}

func TestRelated(t *testing.T) {
	testutil.TestRelated(t, newParser)
}
