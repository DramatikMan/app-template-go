package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelDebug   = slog.LevelDebug
	LevelInfo    = slog.LevelInfo
	LevelWarning = slog.LevelWarn
	LevelError   = slog.LevelError
	LevelFatal   = slog.Level(12)
)

type customHandler struct {
	out     io.Writer
	options *slog.HandlerOptions
	groups  []string
	attrs   []slog.Attr
}

func newHandler(out io.Writer, options *slog.HandlerOptions) *customHandler {
	handler := &customHandler{
		out:     out,
		options: options,
		groups:  []string{},
		attrs:   []slog.Attr{},
	}

	return handler
}

func (handler *customHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= handler.options.Level.Level()
}

func formatTime(value int) string {
	var result string

	if value < 10 {
		result = fmt.Sprintf("0%d", value)
	} else {
		result = fmt.Sprint(value)
	}

	return result
}

func (handler *customHandler) Handle(ctx context.Context, record slog.Record) error {
	levelText, levelName := record.Level.String(), record.Level.String()

	switch record.Level {
	case LevelTrace:
		levelName = "TRACE"
		levelText = fmt.Sprintf("\033[30m\033[44m %s \033[0m", levelName)
	case LevelDebug:
		levelText = fmt.Sprintf("\033[30m\033[46m %s \033[0m", levelText)
	case LevelInfo:
		levelText = fmt.Sprintf("\033[30m\033[42m %s \033[0m", levelText)
	case LevelWarning:
		levelText = fmt.Sprintf("\033[30m\033[43m %s \033[0m", levelText)
	case LevelError:
		levelText = fmt.Sprintf("\033[30m\033[41m %s \033[0m", levelText)
	case LevelFatal:
		levelName = "FATAL"
		levelText = fmt.Sprintf("\033[30m\033[101m %s \033[0m", levelName)
	}

	separator := strings.Repeat(" ", 8-len(levelName))

	msg := fmt.Sprintf(
		"%d/%d/%d %s:%s:%s %s:%s%s\n",
		record.Time.Year(), record.Time.Month(), record.Time.Day(),
		formatTime(record.Time.Hour()),
		formatTime(record.Time.Minute()),
		formatTime(record.Time.Second()),
		levelText, separator, record.Message,
	)

	_, err := handler.out.Write([]byte(msg))

	return err
}

func (handler *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return handler
	}

	newHandler := *handler
	newHandler.attrs = make([]slog.Attr, len(handler.attrs)+len(attrs))
	copy(newHandler.attrs, handler.attrs)

	for i := range attrs {
		newHandler.attrs[len(handler.attrs)+i] = attrs[i]
	}

	return &newHandler
}

func (handler *customHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return handler
	}

	newHandler := *handler
	newHandler.groups = append(handler.groups, name)

	return &newHandler
}
