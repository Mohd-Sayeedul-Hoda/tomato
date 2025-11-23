package demon

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func GetSocketPath() (string, error) {

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
