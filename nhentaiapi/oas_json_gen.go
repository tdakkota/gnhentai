// Code generated by ogen, DO NOT EDIT.

package nhentaiapi

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Book) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Book) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		s.ID.Encode(e)
	}
	{
		e.FieldStart("media_id")
		e.Str(s.MediaID)
	}
	{
		e.FieldStart("images")
		s.Images.Encode(e)
	}
	{
		e.FieldStart("title")
		s.Title.Encode(e)
	}
	{
		e.FieldStart("tags")
		e.ArrStart()
		for _, elem := range s.Tags {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		if s.Scanlator.Set {
			e.FieldStart("scanlator")
			s.Scanlator.Encode(e)
		}
	}
	{
		if s.UploadDate.Set {
			e.FieldStart("upload_date")
			s.UploadDate.Encode(e, json.EncodeUnixSeconds)
		}
	}
	{
		if s.NumPages.Set {
			e.FieldStart("num_pages")
			s.NumPages.Encode(e)
		}
	}
	{
		if s.NumFavorites.Set {
			e.FieldStart("num_favorites")
			s.NumFavorites.Encode(e)
		}
	}
}

var jsonFieldsNameOfBook = [9]string{
	0: "id",
	1: "media_id",
	2: "images",
	3: "title",
	4: "tags",
	5: "scanlator",
	6: "upload_date",
	7: "num_pages",
	8: "num_favorites",
}

// Decode decodes Book from json.
func (s *Book) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Book to nil")
	}
	var requiredBitSet [2]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "media_id":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.MediaID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"media_id\"")
			}
		case "images":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Images.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"images\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "tags":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				s.Tags = make([]Tag, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Tag
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Tags = append(s.Tags, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tags\"")
			}
		case "scanlator":
			if err := func() error {
				s.Scanlator.Reset()
				if err := s.Scanlator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"scanlator\"")
			}
		case "upload_date":
			if err := func() error {
				s.UploadDate.Reset()
				if err := s.UploadDate.Decode(d, json.DecodeUnixSeconds); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"upload_date\"")
			}
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_pages\"")
			}
		case "num_favorites":
			if err := func() error {
				s.NumFavorites.Reset()
				if err := s.NumFavorites.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_favorites\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Book")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b00011111,
		0b00000000,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfBook) {
					name = jsonFieldsNameOfBook[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Book) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Book) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BookID as json.
func (s BookID) Encode(e *jx.Encoder) {
	switch s.Type {
	case IntBookID:
		e.Int(s.Int)
	case StringBookID:
		e.Str(s.String)
	}
}

// Decode decodes BookID from json.
func (s *BookID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BookID to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Number:
		v, err := d.Int()
		s.Int = int(v)
		if err != nil {
			return err
		}
		s.Type = IntBookID
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringBookID
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BookID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BookID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("error")
		s.Error.Encode(e)
	}
}

var jsonFieldsNameOfError = [1]string{
	0: "error",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ErrorDetails as json.
func (s ErrorDetails) Encode(e *jx.Encoder) {
	switch s.Type {
	case StringErrorDetails:
		e.Str(s.String)
	case BoolErrorDetails:
		e.Bool(s.Bool)
	}
}

// Decode decodes ErrorDetails from json.
func (s *ErrorDetails) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorDetails to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Bool:
		v, err := d.Bool()
		s.Bool = bool(v)
		if err != nil {
			return err
		}
		s.Type = BoolErrorDetails
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringErrorDetails
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ErrorDetails) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorDetails) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Image) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Image) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("t")
		e.Str(s.T)
	}
	{
		if s.W.Set {
			e.FieldStart("w")
			s.W.Encode(e)
		}
	}
	{
		if s.H.Set {
			e.FieldStart("h")
			s.H.Encode(e)
		}
	}
}

var jsonFieldsNameOfImage = [3]string{
	0: "t",
	1: "w",
	2: "h",
}

// Decode decodes Image from json.
func (s *Image) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Image to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "t":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.T = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"t\"")
			}
		case "w":
			if err := func() error {
				s.W.Reset()
				if err := s.W.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"w\"")
			}
		case "h":
			if err := func() error {
				s.H.Reset()
				if err := s.H.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"h\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Image")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfImage) {
					name = jsonFieldsNameOfImage[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Image) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Image) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Images) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Images) encodeFields(e *jx.Encoder) {
	{
		if s.Pages != nil {
			e.FieldStart("pages")
			e.ArrStart()
			for _, elem := range s.Pages {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.Cover.Set {
			e.FieldStart("cover")
			s.Cover.Encode(e)
		}
	}
	{
		if s.Thumbnail.Set {
			e.FieldStart("thumbnail")
			s.Thumbnail.Encode(e)
		}
	}
}

var jsonFieldsNameOfImages = [3]string{
	0: "pages",
	1: "cover",
	2: "thumbnail",
}

// Decode decodes Images from json.
func (s *Images) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Images to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "pages":
			if err := func() error {
				s.Pages = make([]Image, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Image
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Pages = append(s.Pages, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pages\"")
			}
		case "cover":
			if err := func() error {
				s.Cover.Reset()
				if err := s.Cover.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover\"")
			}
		case "thumbnail":
			if err := func() error {
				s.Thumbnail.Reset()
				if err := s.Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thumbnail\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Images")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Images) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Images) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Image as json.
func (o OptImage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Image from json.
func (o *OptImage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptImage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptImage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptImage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptNilString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptNilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilString to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptUnixSeconds) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptUnixSeconds) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUnixSeconds to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUnixSeconds) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeUnixSeconds)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUnixSeconds) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeUnixSeconds)
}

// Encode implements json.Marshaler.
func (s *SearchResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SearchResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.ArrStart()
		for _, elem := range s.Result {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		if s.NumPages.Set {
			e.FieldStart("num_pages")
			s.NumPages.Encode(e)
		}
	}
	{
		if s.PerPage.Set {
			e.FieldStart("per_page")
			s.PerPage.Encode(e)
		}
	}
}

var jsonFieldsNameOfSearchResponse = [3]string{
	0: "result",
	1: "num_pages",
	2: "per_page",
}

// Decode decodes SearchResponse from json.
func (s *SearchResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Result = make([]Book, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Book
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Result = append(s.Result, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_pages\"")
			}
		case "per_page":
			if err := func() error {
				s.PerPage.Reset()
				if err := s.PerPage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"per_page\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSearchResponse) {
					name = jsonFieldsNameOfSearchResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SearchResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Tag) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Tag) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int(s.ID)
	}
	{
		e.FieldStart("type")
		s.Type.Encode(e)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		if s.Count.Set {
			e.FieldStart("count")
			s.Count.Encode(e)
		}
	}
}

var jsonFieldsNameOfTag = [5]string{
	0: "id",
	1: "type",
	2: "name",
	3: "url",
	4: "count",
}

// Decode decodes Tag from json.
func (s *Tag) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Tag to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "type":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "count":
			if err := func() error {
				s.Count.Reset()
				if err := s.Count.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Tag")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTag) {
					name = jsonFieldsNameOfTag[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Tag) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Tag) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TagType as json.
func (s TagType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes TagType from json.
func (s *TagType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TagType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch TagType(v) {
	case TagTypeParody:
		*s = TagTypeParody
	case TagTypeCharacter:
		*s = TagTypeCharacter
	case TagTypeTag:
		*s = TagTypeTag
	case TagTypeArtist:
		*s = TagTypeArtist
	case TagTypeGroup:
		*s = TagTypeGroup
	case TagTypeCategory:
		*s = TagTypeCategory
	case TagTypeLanguage:
		*s = TagTypeLanguage
	default:
		*s = TagType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s TagType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TagType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Title) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Title) encodeFields(e *jx.Encoder) {
	{
		if s.English.Set {
			e.FieldStart("english")
			s.English.Encode(e)
		}
	}
	{
		if s.Japanese.Set {
			e.FieldStart("japanese")
			s.Japanese.Encode(e)
		}
	}
	{
		if s.Pretty.Set {
			e.FieldStart("pretty")
			s.Pretty.Encode(e)
		}
	}
}

var jsonFieldsNameOfTitle = [3]string{
	0: "english",
	1: "japanese",
	2: "pretty",
}

// Decode decodes Title from json.
func (s *Title) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Title to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "english":
			if err := func() error {
				s.English.Reset()
				if err := s.English.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"english\"")
			}
		case "japanese":
			if err := func() error {
				s.Japanese.Reset()
				if err := s.Japanese.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"japanese\"")
			}
		case "pretty":
			if err := func() error {
				s.Pretty.Reset()
				if err := s.Pretty.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pretty\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Title")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Title) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Title) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
