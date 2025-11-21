package demon

import (
	"encoding/json"
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

	recoverPanic(manageConn(conn))(conn)
}

func manageConn(conn net.Conn) HandleConn {
	return HandleConn(func(conn net.Conn) {
		conn.Write([]byte("hello world!"))
	})
}
