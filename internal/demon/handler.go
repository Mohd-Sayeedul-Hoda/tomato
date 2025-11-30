package demon

import (
	"context"
	"net"

	"bytes"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
	repo "github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo"
)

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

func createSession(conn net.Conn, sessRepo repo.SessionRepository, req Request) {
	var payload createSessionReq
	err := Decode(bytes.NewReader(req.Data), &payload)
	if err != nil {
		badRequestResponse(conn, req.Method, err)
		return
	}

	estimate := int64(payload.Estimate)
	session := models.Session{
		Label:           payload.Label,
		Note:            &payload.Note,
		IsTracked:       &payload.Tracked,
		SessionEstimate: &estimate,
		Status:          "created",
	}

	id, err := sessRepo.CreateSession(context.Background(), session)
	if err != nil {
		ServerErrorResponse(conn, req.Method, err)
		return
	}

	session.ID = id
	respondWithJSON(conn, req.Method, 200, envelope{"status": "success", "session": session})
}

func startTimer(conn net.Conn, cycleRepo repo.SessionCycleRepository, req Request) {

	ctx := context.Background()
	var req startTimerReq
	err := Decode(bytes.NewReader(req.Data), &req)
	if err != nil {
		badRequestResponse(conn, req.Method, err)
		return
	}

	cycleRepo.ListSessionCycles(ctx, models.SessionCycleFilter{SessionID: reqp})

}
