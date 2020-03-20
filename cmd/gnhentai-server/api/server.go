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

func NewServer(client gnhentai.Client, downloader gnhentai.Downloader, log zerolog.Logger) *Server {
	return &Server{client: client, downloader: downloader, log: log}
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
		pageID = 0
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
	defer s.onClose(image)

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
	defer s.onClose(image)

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
	defer s.onClose(image)

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
	q := req.URL.Query().Get("q")
	if q == "" {
		s.needQueryError(w)
		return
	}

	pageID, ok := s.getPage(w, req)
	if !ok {
		pageID = 0
	}

	r, err := s.client.Search(q, pageID)
	if err != nil {
		s.log.Error().Err(err).Str("query", q).Msg("error caused while searching")
		s.internalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(r)
	if err != nil {
		s.log.Error().Err(err).Str("query", q).Msg("error caused while marshalling doujinshi list")
		s.internalServerError(w)
		return
	}
}

func (s Server) SearchByTag(w http.ResponseWriter, req *http.Request) {
	tagID, err := strconv.Atoi(req.URL.Query().Get("tag_id"))
	if err != nil {
		s.justError(w)
		return
	}

	pageID, ok := s.getPage(w, req)
	if !ok {
		pageID = 0
	}

	r, err := s.client.SearchByTag(gnhentai.Tag{ID: tagID}, pageID)
	if err != nil {
		s.log.Error().Err(err).Int("tag_id", tagID).Msg("error caused while searching")
		s.internalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(r)
	if err != nil {
		s.log.Error().Err(err).Int("tag_id", tagID).Msg("error caused while marshalling doujinshi list")
		s.internalServerError(w)
		return
	}
}

func (s Server) Related(w http.ResponseWriter, req *http.Request) {
	id, ok := s.getBookID(w, req)
	if !ok {
		return
	}

	r, err := s.client.Related(id)
	if err != nil {
		s.log.Error().Err(err).Int("book_id", id).Msg("error caused while getting related from API")
		s.internalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(r)
	if err != nil {
		s.log.Error().Err(err).Int("book_id", id).Msg("error caused while marshalling doujinshi list")
		s.internalServerError(w)
		return
	}
}
