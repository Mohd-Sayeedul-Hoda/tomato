package demon

import (
	"io"
	"log/slog"
)

type envelope map[string]any

func errorResponse(w io.Writer, method string, statusCode int, message any) {
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
	errorResponse(w, method, 500, message)
}

func respondWithJSON[T any](w io.Writer, method string, status int, data T) {

	err := encodeJson(w, data)
	if err != nil {
		ServerErrorResponse(w, method, err)
		return
	}
	slog.Info("api info", slog.String("method", method), slog.Int("status_code", status))
}
