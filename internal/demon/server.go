package demon

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/db"
	repo "github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo"
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo/sqlite"
)

func Serve() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, err := db.OpenSqliteConnection()
	if err != nil {
		return err
	}

	sessionRepo, err := sqlite.NewSessionRepository(db)
	if err != nil {
		return err
	}

	sessCycleRepo, err := sqlite.NewSessionCycleRepository(db)
	if err != nil {
		return err
	}

	socketPath, err := GetSocketPath()
	if err != nil {
		return err
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return err
	}
	defer listener.Close()
	slog.Info("server started")

	var wg sync.WaitGroup
	serverErrors := make(chan error, 1)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				serverErrors <- fmt.Errorf("accept error: %w", err)
				return
			}
			wg.Add(1)
			go handleConn(conn, &wg, sessionRepo, sessCycleRepo)
		}
	}()

	select {
	case <-ctx.Done():
		slog.Info("shutdown signal received")
	case err := <-serverErrors:
		return err
	}

	listener.Close()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		slog.Info("graceful shutdown completed")
	case <-time.After(5 * time.Second):
		slog.Error("shutdown timeout")
	}

	os.Remove(socketPath)
	return nil
}

func handleConn(conn net.Conn, wg *sync.WaitGroup, sessionRepo repo.SessionRepository, sessCycleRepo repo.SessionCycleRepository) {
	defer wg.Done()
	defer conn.Close()

	defer func() {
		if err := recover(); err != nil {
			ServerErrorResponse(conn, NotDefine, fmt.Errorf("%s", err))
			conn.Close()
		}
	}()
	manageConn(conn, sessionRepo, sessCycleRepo)
}

func manageConn(conn net.Conn, sessionRepo repo.SessionRepository, sessCycleRepo repo.SessionCycleRepository) {

	var req Request
	err := Decode(conn, &req)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidRequest):
			errorResponse(conn, NotDefine, 400, err.Error())
		default:
			ServerErrorResponse(conn, NotDefine, err)
		}
		return
	}

	switch req.Method {
	case Status:
		healthCheck(conn, req)
	case SessionCreate:
		createSession(conn, sessionRepo, req)
	default:
		notFound(conn)
	}
}
