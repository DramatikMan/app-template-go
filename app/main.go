package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	fmt.Println("Go environment.")

	handler := newHandler(os.Stdout, &slog.HandlerOptions{Level: LevelTrace})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Log(context.Background(), LevelTrace, "trace log")
	slog.Debug("debug log")
	slog.Info("info log")
	slog.Warn("warning log")
	slog.Error("error log")
	slog.Log(context.Background(), LevelFatal, "fatal error log")
}
