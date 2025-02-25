// Code generated by ogen, DO NOT EDIT.

package nhentaiapi

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

type codeRecorder struct {
	http.ResponseWriter
	status int
}

func (c *codeRecorder) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func recordError(string, error) {}

// handleGetBookRequest handles getBook operation.
//
// Gets metadata of book.
//
// GET /api/gallery/{book_id}
func (s *Server) handleGetBookRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetBookOperation,
			ID:   "getBook",
		}
	)
	params, err := decodeGetBookParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Book
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetBookOperation,
			OperationSummary: "Gets metadata of book",
			OperationID:      "getBook",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "book_id",
					In:   "path",
				}: params.BookID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetBookParams
			Response = *Book
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetBookParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetBook(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetBook(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeGetBookResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPageCoverImageRequest handles getPageCoverImage operation.
//
// Gets page cover.
//
// GET /galleries/{media_id}/cover.{format}
func (s *Server) handleGetPageCoverImageRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPageCoverImageOperation,
			ID:   "getPageCoverImage",
		}
	)
	params, err := decodeGetPageCoverImageParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *GetPageCoverImageOKHeaders
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPageCoverImageOperation,
			OperationSummary: "Gets page cover",
			OperationID:      "getPageCoverImage",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "media_id",
					In:   "path",
				}: params.MediaID,
				{
					Name: "format",
					In:   "path",
				}: params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPageCoverImageParams
			Response = *GetPageCoverImageOKHeaders
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPageCoverImageParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPageCoverImage(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPageCoverImage(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeGetPageCoverImageResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPageImageRequest handles getPageImage operation.
//
// Gets page.
//
// GET /galleries/{media_id}/{page}.{format}
func (s *Server) handleGetPageImageRequest(args [3]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPageImageOperation,
			ID:   "getPageImage",
		}
	)
	params, err := decodeGetPageImageParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *GetPageImageOKHeaders
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPageImageOperation,
			OperationSummary: "Gets page",
			OperationID:      "getPageImage",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "media_id",
					In:   "path",
				}: params.MediaID,
				{
					Name: "page",
					In:   "path",
				}: params.Page,
				{
					Name: "format",
					In:   "path",
				}: params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPageImageParams
			Response = *GetPageImageOKHeaders
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPageImageParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPageImage(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPageImage(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeGetPageImageResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetPageThumbnailImageRequest handles getPageThumbnailImage operation.
//
// Gets page thumbnail.
//
// GET /galleries/{media_id}/{page}t.{format}
func (s *Server) handleGetPageThumbnailImageRequest(args [3]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: GetPageThumbnailImageOperation,
			ID:   "getPageThumbnailImage",
		}
	)
	params, err := decodeGetPageThumbnailImageParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *GetPageThumbnailImageOKHeaders
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    GetPageThumbnailImageOperation,
			OperationSummary: "Gets page thumbnail",
			OperationID:      "getPageThumbnailImage",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "media_id",
					In:   "path",
				}: params.MediaID,
				{
					Name: "page",
					In:   "path",
				}: params.Page,
				{
					Name: "format",
					In:   "path",
				}: params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPageThumbnailImageParams
			Response = *GetPageThumbnailImageOKHeaders
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPageThumbnailImageParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetPageThumbnailImage(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetPageThumbnailImage(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeGetPageThumbnailImageResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleRelatedRequest handles related operation.
//
// Search for related comics.
//
// GET /api/galleries/{book_id}/related
func (s *Server) handleRelatedRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: RelatedOperation,
			ID:   "related",
		}
	)
	params, err := decodeRelatedParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *SearchResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    RelatedOperation,
			OperationSummary: "Search for related comics",
			OperationID:      "related",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "book_id",
					In:   "path",
				}: params.BookID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RelatedParams
			Response = *SearchResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRelatedParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Related(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Related(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeRelatedResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSearchRequest handles search operation.
//
// Search for comics.
//
// GET /api/galleries/search
func (s *Server) handleSearchRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SearchOperation,
			ID:   "search",
		}
	)
	params, err := decodeSearchParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *SearchResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SearchOperation,
			OperationSummary: "Search for comics",
			OperationID:      "search",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "query",
					In:   "query",
				}: params.Query,
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "per_page",
					In:   "query",
				}: params.PerPage,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchParams
			Response = *SearchResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Search(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Search(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeSearchResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSearchByTagIDRequest handles searchByTagID operation.
//
// Search for comics by tag ID.
//
// GET /api/galleries/tagged
func (s *Server) handleSearchByTagIDRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: SearchByTagIDOperation,
			ID:   "searchByTagID",
		}
	)
	params, err := decodeSearchByTagIDParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *SearchResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    SearchByTagIDOperation,
			OperationSummary: "Search for comics by tag ID",
			OperationID:      "searchByTagID",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "tag_id",
					In:   "query",
				}: params.TagID,
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "per_page",
					In:   "query",
				}: params.PerPage,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchByTagIDParams
			Response = *SearchResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchByTagIDParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SearchByTagID(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.SearchByTagID(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeSearchByTagIDResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
