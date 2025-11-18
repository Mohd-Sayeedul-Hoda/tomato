/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log/slog"
	"os"

	"github.com/Mohd-Sayeedul-Hoda/tomato/cmd"
)

func main() {
	slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	cmd.Execute()
}
