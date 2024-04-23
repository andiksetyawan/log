package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/andiksetyawan/log"
)

type Log struct {
	logger *slog.Logger
	level  Level
}

type OptFunc func(*Log) error

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func WithLogger(logger *slog.Logger) OptFunc {
	return func(sl *Log) (err error) {
		sl.logger = logger
		return
	}
}

func WithLevelString(s string) OptFunc {
	var l Level = -1
	switch strings.ToLower(s) {
	case "debug":
		l = 0
	case "info":
		l = 1
	case "warn":
		l = 2
	case "error":
		l = 3
	case "fatal":
		l = 4
	}
	return WithLevel(l)
}

func WithLevel(l Level) OptFunc {
	return func(sl *Log) (err error) {
		if l < LevelDebug || l > LevelFatal {
			return fmt.Errorf("invalid level: %d", l)
		}

		sl.level = l
		return
	}
}

func New(opts ...OptFunc) (l log.Logger, err error) {
	s := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	sl := &Log{logger: s, level: LevelInfo}
	for _, opt := range opts {
		err = opt(sl)
		if err != nil {
			return nil, fmt.Errorf("fail to apply options: %w", err)
		}
	}
	return sl, nil
}

func (l Log) Debug(ctx context.Context, message string, args ...any) {
	if l.level > LevelDebug {
		return
	}

	l.logger.DebugContext(ctx, message, args...)
}

func (l Log) Info(ctx context.Context, message string, args ...any) {
	if l.level > LevelInfo {
		return
	}

	l.logger.InfoContext(ctx, message, args...)
}

func (l Log) Warn(ctx context.Context, message string, args ...any) {
	if l.level > LevelWarn {
		return
	}

	l.logger.WarnContext(ctx, message, args...)
}

func (l Log) Error(ctx context.Context, message string, args ...any) {
	if l.level > LevelError {
		return
	}

	l.logger.ErrorContext(ctx, message, args...)
}
