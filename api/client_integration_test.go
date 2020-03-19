package api

import (
	"github.com/tdakkota/gnhentai"
	"github.com/tdakkota/gnhentai/testutil"
	"testing"
)

func newClient(t *testing.T) gnhentai.Client {
	return NewClient(WithClient(testutil.TestClient(t)))
}

func TestRandom(t *testing.T) {
	testutil.TestRandom(t, newClient)
}

func TestSearch(t *testing.T) {
	testutil.TestSearch(t, newClient)
}

func TestSearchByTag(t *testing.T) {
	testutil.TestSearchByTag(t, newClient)
}

func TestGetByID(t *testing.T) {
	testutil.TestGetByID(t, newClient)
}

func TestGetByID2(t *testing.T) {
	testutil.TestGetByID2(t, newClient)
}

func TestRelated(t *testing.T) {
	testutil.TestRelated(t, newClient)
}
