package api

import (
	_ "embed"
	"testing"

	"github.com/go-faster/jx"
	"github.com/stretchr/testify/require"
	"github.com/tdakkota/gnhentai/internal/nhentaiapi"
)

//go:embed _testdata/search.json
var search []byte

func TestSearchResponse(t *testing.T) {
	checkSchema[nhentaiapi.SearchResponse](t, search)
}

//go:embed _testdata/book.json
var book []byte

func TestBook(t *testing.T) {
	checkSchema[nhentaiapi.Book](t, book)
}

func checkSchema[
	T any,
	P interface {
		*T
		Decode(*jx.Decoder) error
		Encode(*jx.Encoder)
	},
](t *testing.T, data []byte) {
	var b T
	require.NoError(t, P(&b).Decode(jx.DecodeBytes(data)))

	e := jx.GetEncoder()
	P(&b).Encode(e)

	var b2 T
	require.NoError(t, P(&b2).Decode(jx.DecodeBytes(e.Bytes())))
	require.Equal(t, b, b2)

	require.JSONEq(t, string(data), e.String())
}
