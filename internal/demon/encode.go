package demon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	Method string          `json:"method"`
	Data   json.RawMessage `json:"data"`
}

type envelope map[string]any

var ErrInvalidData = errors.New("invalid data")
var ErrInvalidRequest = errors.New("invalid request")

func encodeJson[T any](w io.Writer, data T) error {

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	err := encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func Decode[T any](r io.Reader, data *T) error {

	maxBytes := 1_048_576
	r = io.LimitReader(r, int64(maxBytes))
	dec := json.NewDecoder(r)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("%w: body contain badly-formated JSON (at character %d)", ErrInvalidRequest, syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("%w: body contains badly-formated JSON", ErrInvalidRequest)

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("%w: body contains incorrect JSON type for field %q", ErrInvalidRequest, unmarshalTypeError.Field)
			}
			return fmt.Errorf("%w: body contians incorrect JSON type (at character %d)", ErrInvalidRequest, unmarshalTypeError.Offset)

		// happen when body is empty
		case errors.Is(err, io.EOF):
			return fmt.Errorf("%w: body must not be empty", ErrInvalidRequest)

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("%w: body contains unknown key %s", ErrInvalidRequest, fieldName)

			// invalid argument when pass to decode function
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	return nil
}
