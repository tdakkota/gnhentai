// Code generated by ogen, DO NOT EDIT.

package nhentaiapi

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetBookParams is parameters of getBook operation.
type GetBookParams struct {
	// ID of book.
	BookID string
}

func unpackGetBookParams(packed middleware.Parameters) (params GetBookParams) {
	{
		key := middleware.ParameterKey{
			Name: "book_id",
			In:   "path",
		}
		params.BookID = packed[key].(string)
	}
	return params
}

func decodeGetBookParams(args [1]string, argsEscaped bool, r *http.Request) (params GetBookParams, _ error) {
	// Decode path: book_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "book_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.BookID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "book_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetPageCoverImageParams is parameters of getPageCoverImage operation.
type GetPageCoverImageParams struct {
	// ID of book.
	MediaID string
	// Image format.
	Format string
}

func unpackGetPageCoverImageParams(packed middleware.Parameters) (params GetPageCoverImageParams) {
	{
		key := middleware.ParameterKey{
			Name: "media_id",
			In:   "path",
		}
		params.MediaID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "format",
			In:   "path",
		}
		params.Format = packed[key].(string)
	}
	return params
}

func decodeGetPageCoverImageParams(args [2]string, argsEscaped bool, r *http.Request) (params GetPageCoverImageParams, _ error) {
	// Decode path: media_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "media_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MediaID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "media_id",
			In:   "path",
			Err:  err,
		}
	}
	// Set default value for path: format.
	{
		val := string("jpg")
		params.Format = val
	}
	// Decode path: format.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "format",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Format = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "format",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetPageImageParams is parameters of getPageImage operation.
type GetPageImageParams struct {
	// ID of book.
	MediaID string
	// Number of page.
	Page int
	// Image format.
	Format string
}

func unpackGetPageImageParams(packed middleware.Parameters) (params GetPageImageParams) {
	{
		key := middleware.ParameterKey{
			Name: "media_id",
			In:   "path",
		}
		params.MediaID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "path",
		}
		params.Page = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "format",
			In:   "path",
		}
		params.Format = packed[key].(string)
	}
	return params
}

func decodeGetPageImageParams(args [3]string, argsEscaped bool, r *http.Request) (params GetPageImageParams, _ error) {
	// Decode path: media_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "media_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MediaID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "media_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: page.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "page",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.Page = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           0,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.Page)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "path",
			Err:  err,
		}
	}
	// Set default value for path: format.
	{
		val := string("jpg")
		params.Format = val
	}
	// Decode path: format.
	if err := func() error {
		param := args[2]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[2])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "format",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Format = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "format",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetPageThumbnailImageParams is parameters of getPageThumbnailImage operation.
type GetPageThumbnailImageParams struct {
	// ID of book.
	MediaID string
	// Number of page.
	Page int
	// Image format.
	Format string
}

func unpackGetPageThumbnailImageParams(packed middleware.Parameters) (params GetPageThumbnailImageParams) {
	{
		key := middleware.ParameterKey{
			Name: "media_id",
			In:   "path",
		}
		params.MediaID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "path",
		}
		params.Page = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "format",
			In:   "path",
		}
		params.Format = packed[key].(string)
	}
	return params
}

func decodeGetPageThumbnailImageParams(args [3]string, argsEscaped bool, r *http.Request) (params GetPageThumbnailImageParams, _ error) {
	// Decode path: media_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "media_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MediaID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "media_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: page.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "page",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.Page = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           0,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.Page)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "path",
			Err:  err,
		}
	}
	// Set default value for path: format.
	{
		val := string("jpg")
		params.Format = val
	}
	// Decode path: format.
	if err := func() error {
		param := args[2]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[2])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "format",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Format = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "format",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RelatedParams is parameters of related operation.
type RelatedParams struct {
	// ID of book.
	BookID string
}

func unpackRelatedParams(packed middleware.Parameters) (params RelatedParams) {
	{
		key := middleware.ParameterKey{
			Name: "book_id",
			In:   "path",
		}
		params.BookID = packed[key].(string)
	}
	return params
}

func decodeRelatedParams(args [1]string, argsEscaped bool, r *http.Request) (params RelatedParams, _ error) {
	// Decode path: book_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "book_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.BookID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "book_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// SearchParams is parameters of search operation.
type SearchParams struct {
	// Search query.
	// * You can search for multiple terms at the same time, and this will return only galleries that
	// contain both terms.
	// For example, `anal tanlines` finds all galleries that contain both `anal` and `tanlines`.
	// * You can exclude terms by prefixing them with `-`. For example, `anal tanlines -yaoi` matches all
	// galleries matching `anal` and `tanlines` but not `yaoi`.
	// * Exact searches can be performed by wrapping terms in double quotes. For example, `"big breasts"`
	// only matches galleries with `"big breasts"` exactly somewhere in the title or in tags.
	// * These can be combined with tag namespaces for finer control over the query: `parodies:railgun
	// -tag:"big breasts"`.
	Query string
	// Number of result page.
	Page OptInt
}

func unpackSearchParams(packed middleware.Parameters) (params SearchParams) {
	{
		key := middleware.ParameterKey{
			Name: "query",
			In:   "query",
		}
		params.Query = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	return params
}

func decodeSearchParams(args [0]string, argsEscaped bool, r *http.Request) (params SearchParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: query.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "query",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Query = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "query",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SearchByTagIDParams is parameters of searchByTagID operation.
type SearchByTagIDParams struct {
	// Tag ID.
	TagID int
	// Number of result page.
	Page OptInt
}

func unpackSearchByTagIDParams(packed middleware.Parameters) (params SearchByTagIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "tag_id",
			In:   "query",
		}
		params.TagID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	return params
}

func decodeSearchByTagIDParams(args [0]string, argsEscaped bool, r *http.Request) (params SearchByTagIDParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: tag_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "tag_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.TagID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "tag_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
