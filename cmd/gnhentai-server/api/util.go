package api

import (
	"io"
	"net/http"
	"strconv"
)

func getParam(name string, req *http.Request) (string, bool) {
	return "", false
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

func (s Server) getIntParam(name string, w http.ResponseWriter, req *http.Request) (id int, ok bool) {
	var err error

	if v, ok := getParam(name, req); ok {
		id, err = strconv.Atoi(v)
		if err != nil || id <= 0 {
			s.justError(w)
			return
		}
	}

	return id, true
}

func (s Server) getBookID(w http.ResponseWriter, req *http.Request) (id int, ok bool) {
	return s.getIntParam("book_id", w, req)
}

func (s Server) getPage(w http.ResponseWriter, req *http.Request) (id int, ok bool) {
	return s.getIntParam("page", w, req)
}

func (s Server) onClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		s.log.Error().Err(err).Msg("close failed")
	}
}
