package application

import (
	"context"
	"fmt"
	"log/slog"
)

type level = slog.Level

const levelDryRun level = slog.Level(-1)

type logger struct{ *slog.Logger }

func (l *logger) DryRun(msg string, args ...any) {
	ctx := context.Background()
	l.Log(ctx, levelDryRun, msg, args...)
}

func (c *Compose) DryRunMessage(mlev int, msg string) {
	if c.DryRun {
		var prefix string
		ansiReset := "\033[0m"
		ansiLightBlue := "\u001b[36m"

		if mlev == 1 {
			prefix = " =>"
		}
		val := fmt.Sprintf("%s%s%s%s", ansiLightBlue, "DRY-RUN =>", prefix, ansiReset)
		fmt.Printf("%s %s\n", val, msg)
	}
}
