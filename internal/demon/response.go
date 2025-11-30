package demon

import (
	"io"
	"log/slog"
	"net"
)

type envelope map[string]any

func errorResponse(w io.Writer, method Method, statusCode int, message any) {
	resp := envelope{
		"status": "failed",
		"error":  message,
	}

	err := encodeJson(w, resp)
	if err != nil {
		slog.Error("server error", slog.String("err", err.Error()), slog.Any("method", method))
		return
	}
}

func ServerErrorResponse(w io.Writer, method Method, err error) {
	slog.Error("server error", slog.Any("method", method), slog.String("err", err.Error()))
	message := "Internal Server Error"
	errorResponse(w, method, 500, message)
}

func badRequestResponse(w net.Conn, method Method, err error) {
	errorResponse(w, method, 400, err.Error())
}

func respondWithJSON[T any](w io.Writer, method Method, status int, data T) {

	err := encodeJson(w, data)
	if err != nil {
		ServerErrorResponse(w, method, err)
		return
	}
	slog.Info("api info", slog.Any("method", method), slog.Int("status_code", status))
}
