package server

import (
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"strconv"
)

func getParam(name string, req *http.Request) (string, bool) {
	return chi.URLParam(req, name), name != ""
}

func (s Server) justError(w http.ResponseWriter) {
	_, _ = w.Write([]byte(`{"error": true}`))
}

func (s Server) internalServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	s.justError(w)
}

func (s Server) needQueryError(w http.ResponseWriter) {
	w.WriteHeader(403)
	_, _ = w.Write([]byte(`{"error": "You need to provide a search query"}`))
}

func (s Server) getIntParam(name string, req *http.Request) (id int, ok bool) {
	var err error

	if v, ok := getParam(name, req); ok {
		id, err = strconv.Atoi(v)
		if err != nil || id <= 0 {
			return 0, false
		}
		return id, true
	}

	return id, false
}

func (s Server) getBookID(req *http.Request) (id int, ok bool) {
	return s.getIntParam("book_id", req)
}

func (s Server) getPage(req *http.Request) (id int, ok bool) {
	return s.getIntParam("page", req)
}

func (s Server) onClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		s.log.Error().Err(err).Msg("close failed")
	}
}
