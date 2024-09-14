package application

import (
	"context"
	"log/slog"
)

type level = slog.Level

const levelDryRun level = slog.Level(-1)

type logger struct{ *slog.Logger }

func (l *logger) DryRun(msg string, args ...any) {
	ctx := context.Background()
	l.Log(ctx, levelDryRun, msg, args...)
}
