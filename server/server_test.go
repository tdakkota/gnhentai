package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/tdakkota/gnhentai"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

type MockClient struct {
	test gnhentai.Doujinshi
}

func (m MockClient) Test() []byte {
	data, _ := json.Marshal(m.test)
	return data
}

func (m MockClient) TestString() string {
	return string(m.Test())
}

func (m MockClient) Page(mediaID, n int, format string) (io.ReadCloser, error) {
	if mediaID == 13 {
		return nil, errors.New("i'm superstitious")
	}
	return ioutil.NopCloser(bytes.NewBufferString("pretty page")), nil
}

func (m MockClient) Thumbnail(mediaID int, n int, format string) (io.ReadCloser, error) {
	if mediaID == 13 {
		return nil, errors.New("i'm superstitious")
	}
	return ioutil.NopCloser(bytes.NewBufferString("pretty thumbnail")), nil
}

func (m MockClient) Cover(mediaID int, format string) (io.ReadCloser, error) {
	if mediaID == 13 {
		return nil, errors.New("i'm superstitious")
	}
	return ioutil.NopCloser(bytes.NewBufferString("pretty cover")), nil
}

func (m MockClient) ByID(id int) (gnhentai.Doujinshi, error) {
	if id == 13 {
		return m.test, errors.New("i'm superstitious")
	}
	return m.test, nil
}

func (m MockClient) Random() (gnhentai.Doujinshi, error) {
	return m.test, nil
}

func (m MockClient) Search(q string, page int) ([]gnhentai.Doujinshi, error) {
	return []gnhentai.Doujinshi{m.test}, nil
}

func (m MockClient) SearchByTag(tag gnhentai.Tag, page int) ([]gnhentai.Doujinshi, error) {
	return []gnhentai.Doujinshi{m.test}, nil
}

func (m MockClient) Related(id int) ([]gnhentai.Doujinshi, error) {
	return []gnhentai.Doujinshi{m.test}, nil
}

func testServer() (*Server, MockClient, error) {
	m := MockClient{}
	data, err := ioutil.ReadFile("../testdata/305329.json")
	if err != nil {
		return nil, m, err
	}

	err = json.Unmarshal(data, &m.test)
	if err != nil {
		return nil, m, err
	}

	return NewServer(m, m, WithLogger(zerolog.New(os.Stdout))), m, nil
}

func chiParams(req *http.Request, params map[string]string) {
	if req != nil {
		if v := req.Context().Value(chi.RouteCtxKey); v == nil {
			routeCtx := chi.NewRouteContext()
			for k, v := range params {
				routeCtx.URLParams.Add(k, v)
			}
			*req = *req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, routeCtx))
		}
	}
}

type handlerArgs struct {
	pathArgs  map[string]string
	queryArgs url.Values
}

func testHandler(t *testing.T, handler http.Handler, args handlerArgs, result string, code int) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	if args.queryArgs != nil {
		req.URL.RawQuery = args.queryArgs.Encode()
	}

	if args.pathArgs != nil {
		chiParams(req, args.pathArgs)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	s := rr.Body.String()
	t.Log(s)
	require.Equal(t, code, rr.Code)
	require.Equal(t, strings.TrimSpace(result), strings.TrimSpace(rr.Body.String()))
}

const bookID = 305329

func TestServer_GetBookByID(t *testing.T) {
	s, m, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetBookByID), handlerArgs{}, `{"error": true}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetBookByID), handlerArgs{
			pathArgs: map[string]string{"book_id": strconv.Itoa(bookID)},
		}, m.TestString(), 200)
	})

	t.Run("ok-page", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetBookByID), handlerArgs{
			pathArgs: map[string]string{
				"book_id": strconv.Itoa(bookID),
				"page":    "2",
			},
		}, m.TestString(), 200)
	})
}

func TestServer_GetCoverByID(t *testing.T) {
	s, _, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetCoverByID), handlerArgs{}, `{"error": true}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetCoverByID), handlerArgs{
			pathArgs: map[string]string{
				"book_id": strconv.Itoa(bookID),
			},
		}, "pretty cover", 200)
	})
}

func TestServer_GetPageByID(t *testing.T) {
	s, _, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetPageByID), handlerArgs{}, `{"error": true}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetPageByID), handlerArgs{
			pathArgs: map[string]string{
				"book_id": strconv.Itoa(bookID),
				"page":    "2",
			},
		}, "pretty page", 200)
	})
}

func TestServer_GetThumbnailByID(t *testing.T) {
	s, _, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetThumbnailByID), handlerArgs{}, `{"error": true}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.GetThumbnailByID), handlerArgs{
			pathArgs: map[string]string{
				"book_id": strconv.Itoa(bookID),
				"page":    "2",
			},
		}, "pretty thumbnail", 200)
	})
}

func TestServer_Related(t *testing.T) {
	s, m, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.Related), handlerArgs{}, `{"error": true}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.Related), handlerArgs{
			pathArgs: map[string]string{
				"book_id": strconv.Itoa(bookID),
			},
			// i'm fuckin array now
		}, "["+m.TestString()+"]", 200)
	})
}

func TestServer_Search(t *testing.T) {
	s, m, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.Search), handlerArgs{}, `{"error": "You need to provide a search query"}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		args := url.Values{}
		args.Set("query", "milf")
		testHandler(t, http.HandlerFunc(s.Search), handlerArgs{
			queryArgs: args,
			// i'm fuckin array now
		}, "["+m.TestString()+"]", 200)
	})
}

func TestServer_SearchByTag(t *testing.T) {
	s, m, err := testServer()
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("error-no-id", func(t *testing.T) {
		testHandler(t, http.HandlerFunc(s.SearchByTag), handlerArgs{}, `{"error": true}`, 403)
	})

	t.Run("ok", func(t *testing.T) {
		args := url.Values{}
		args.Set("tag_id", "69")
		testHandler(t, http.HandlerFunc(s.SearchByTag), handlerArgs{
			queryArgs: args,
			// i'm fuckin array now
		}, "["+m.TestString()+"]", 200)
	})
}
