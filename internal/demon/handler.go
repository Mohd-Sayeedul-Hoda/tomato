package demon

import (
	"encoding/json"
	"net"

	repo "github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo"
)

type Request struct {
	Method string          `json:"method"`
	Data   json.RawMessage `json:"data"`
}

type startTimerReq struct {
	Temp bool `json:"temp"`
}

func healthCheck(conn net.Conn, req Request) {
	resp := envelope{
		"status":  "success",
		"message": "service is running",
	}

	respondWithJSON(conn, req.Method, 200, resp)
}

func notFound(conn net.Conn) {
	resp := envelope{
		"status": "failed",
		"error":  "method not found",
	}

	respondWithJSON(conn, "NOT_FOUND", 400, resp)
}

func startSession(conn net.Conn, sessRepo repo.SessionRepository, sessCycleRepo repo.SessionCycleRepository, req Request) {

}
