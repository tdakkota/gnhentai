package server

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/tdakkota/gnhentai"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	cache      Cache
	client     gnhentai.Client
	downloader gnhentai.Downloader
	log        zerolog.Logger
}

func NewServer(client gnhentai.Client, downloader gnhentai.Downloader, options ...Option) *Server {
	s := &Server{client: client, downloader: downloader}

	for _, opt := range options {
		opt(s)
	}

	return s
}

func (s Server) GetBookByID(w http.ResponseWriter, req *http.Request) {
	id, ok := s.getBookID(req)
	if !ok {
		s.justError(w)
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
	bookID, ok := s.getBookID(req)
	if !ok {
		s.justError(w)
		return
	}

	pageID, ok := s.getPage(req)
	if !ok {
		pageID = 0
	}

	format, ok := getParam("format", req)
	if !ok {
		format = "jpg"
	}

	image, err := s.downloader.Page(bookID, pageID, format)
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
	bookID, ok := s.getBookID(req)
	if !ok {
		s.justError(w)
		return
	}

	image, err := s.downloader.Cover(bookID, "")
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
	bookID, ok := s.getBookID(req)
	if !ok {
		s.justError(w)
		return
	}

	pageID, ok := s.getPage(req)
	if !ok {
		return
	}

	image, err := s.downloader.Thumbnail(bookID, pageID, "")
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
	q := req.URL.Query().Get("query")
	if q == "" {
		s.needQueryError(w)
		return
	}

	pageID, ok := s.getPage(req)
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

	pageID, ok := s.getPage(req)
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
	id, ok := s.getBookID(req)
	if !ok {
		s.justError(w)
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

func (s Server) Register(r chi.Router) {
	r.Route("/api/galleries", func(r chi.Router) {
		r.Get("/search", s.Search)
		r.Get("/tagged", s.SearchByTag)
	})

	r.Route("/api/gallery", func(r chi.Router) {
		r.Get("/{book_id}", s.GetBookByID)
		r.Get("/{book_id}/related", s.Related)
	})

	r.Route("/galleries", func(r chi.Router) {
		r.Get("/{book_id}/{page}.{format}", s.GetPageByID)
		r.Get("/{book_id}/{page}t.{format}", s.GetThumbnailByID)
		r.Get("/{book_id}/cover.{format}", s.GetCoverByID)
	})
}
