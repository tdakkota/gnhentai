package api

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/tdakkota/gnhentai"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	client     gnhentai.Client
	downloader gnhentai.Downloader
	log        zerolog.Logger
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

func (s Server) GetBookByID(w http.ResponseWriter, req *http.Request) {
	id, ok := s.getBookID(w, req)
	if !ok {
		return
	}

	doujinshi, err := s.client.ByID(id)
	if err != nil {
		s.log.Error().Err(err).Int("book_id", id).Msg("error caused while getting book from API")
		s.internalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(doujinshi)
	if err != nil {
		s.log.Error().Err(err).Int("book_id", id).Msg("error caused while marshalling doujinshi")
		s.internalServerError(w)
		return
	}
}

func (s Server) GetPageByID(w http.ResponseWriter, req *http.Request) {
	bookID, ok := s.getBookID(w, req)
	if !ok {
		return
	}

	pageID, ok := s.getPage(w, req)
	if !ok {
		return
	}

	image, err := s.downloader.Page(bookID, pageID)
	if err != nil {
		s.log.Error().Err(err).
			Int("book_id", bookID).
			Int("page_id", pageID).
			Msg("error caused while getting page from API")
		s.internalServerError(w)
		return
	}
	defer image.Close()

	_, err = io.Copy(w, image)
	if err != nil {
		s.log.Error().Err(err).
			Int("book_id", bookID).
			Int("page_id", pageID).
			Msg("error caused in io.Copy")
		s.internalServerError(w)
		return
	}
}

func (s Server) GetCoverByID(w http.ResponseWriter, req *http.Request) {
	bookID, ok := s.getBookID(w, req)
	if !ok {
		return
	}

	image, err := s.downloader.Cover(bookID)
	if err != nil {
		s.log.Error().Err(err).
			Int("book_id", bookID).
			Msg("error caused while getting cover from API")
		s.internalServerError(w)
		return
	}
	defer image.Close()

	_, err = io.Copy(w, image)
	if err != nil {
		s.log.Error().Err(err).
			Int("book_id", bookID).
			Msg("error caused in io.Copy")
		s.internalServerError(w)
		return
	}
}

func (s Server) GetThumbnailByID(w http.ResponseWriter, req *http.Request) {
	bookID, ok := s.getBookID(w, req)
	if !ok {
		return
	}

	pageID, ok := s.getPage(w, req)
	if !ok {
		return
	}

	image, err := s.downloader.Thumbnail(bookID, pageID)
	if err != nil {
		s.log.Error().Err(err).
			Int("book_id", bookID).
			Int("page_id", pageID).
			Msg("error caused while getting thumbnail from API")
		s.internalServerError(w)
		return
	}
	defer image.Close()

	_, err = io.Copy(w, image)
	if err != nil {
		s.log.Error().Err(err).
			Int("book_id", bookID).
			Int("page_id", pageID).
			Msg("error caused in io.Copy")
		s.internalServerError(w)
		return
	}
}

func (s Server) Search(w http.ResponseWriter, req *http.Request) {
}

func (s Server) SearchByTag(w http.ResponseWriter, req *http.Request) {
}

func (s Server) Related(w http.ResponseWriter, req *http.Request) {
}
