package main

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/tdakkota/gnhentai"
	"net/http"
	"strconv"
)

type Server struct {
	client gnhentai.Client
	log    zerolog.Logger
}

func NewServer(client gnhentai.Client, log zerolog.Logger) *Server {
	return &Server{client: client, log: log}
}

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

func (s Server) GetByID(w http.ResponseWriter, req *http.Request) {
	var id int
	var err error

	if v, ok := getParam("book_id", req); ok {
		id, err = strconv.Atoi(v)
		if err != nil || id <= 0 {
			s.justError(w)
			return
		}
	}

	doujinshi, err := s.client.ByID(id)
	if err != nil {
		s.log.Error().Err(err).Int("book_id", id).Msg("error caused while getting book from API")
		s.internalServerError(w)
		return
	}

	err = json.NewEncoder(w).Encode(doujinshi)
	if err != nil {
		s.log.Error().Err(err).Int("book_id", id).Msg("error caused while marshalling doujinshi")
		s.internalServerError(w)
		return
	}
}
