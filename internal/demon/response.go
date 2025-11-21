package demon

import (
	"io"
	"log/slog"
)

type envelope map[string]any

func ErrorResponse(w io.Writer, method string, statusCode int, message any) {
	resp := envelope{
		"status": "failed",
		"error":  message,
	}

	err := encodeJson(w, resp)
	if err != nil {
		slog.Error("server error", slog.String("err", err.Error()), slog.String("method", method))
		return
	}

}

func ServerErrorResponse(w io.Writer, method string, err error) {
	slog.Error("server error", slog.String("method", method), slog.String("err", err.Error()))
	message := "Internal Server Error"
	ErrorResponse(w, method, 500, message)
}
