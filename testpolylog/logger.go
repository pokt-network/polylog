package testpolylog

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/pokt-network/polylog"
	"github.com/pokt-network/polylog/polyzero"
)

func NewLoggerWithCtx(
	ctx context.Context,
	level polylog.Level,
) (polylog.Logger, context.Context) {
	levelOpt := polyzero.WithLevel(zerolog.Level(level.Int()))
	logger := polyzero.NewLogger(levelOpt)
	ctx = logger.WithContext(ctx)

	return logger, ctx
}
