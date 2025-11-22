package demon

import (
	"encoding/json"
	"errors"
	"net"
	"sync"
)

type Request struct {
	Method string          `json:"method"`
	Data   json.RawMessage `json:"data"`
}

func handleConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	recoverPanic(manageConn())(conn)
}

func manageConn() HandleConn {
	return HandleConn(func(conn net.Conn) {

		var req Request
		err := Decode(conn, &req)
		if err != nil {
			switch {
			case errors.Is(err, ErrInvalidRequest):
				errorResponse(conn, "NONE", 400, err.Error())
			default:
				ServerErrorResponse(conn, "NONE", err)
			}
			return
		}

		switch req.Method {
		case "STATUS":
			healthCheck(conn, req)
		default:
			notFound(conn)
		}
	})
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
