package demon

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"
)

func getSocketPath() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	socketDir := filepath.Join(homeDir, ".local", "share", "tomato")
	if err := os.MkdirAll(socketDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create socket directory: %w", err)
	}

	return filepath.Join(socketDir, "tomato.sock"), nil

}

func Serve() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	socketPath, err := getSocketPath()
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
			go handleConn(conn, &wg)
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

func handleConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
}
