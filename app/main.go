package main

import (
	"context"
	"fmt"
	"log/slog"
	"project/app/clog"
)

func main() {
	fmt.Println("Go environment.")

	handler := clog.NewConsoleHandler(&slog.HandlerOptions{Level: clog.LevelTrace})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Log(context.Background(), clog.LevelTrace, "trace log")
	slog.Debug("debug log")
	slog.Info("info log")
	slog.Warn("warning log")
	slog.Error("error log")
	slog.Log(context.Background(), clog.LevelFatal, "fatal error log")
}
