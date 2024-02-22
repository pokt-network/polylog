package polystd

import (
	"io"
	"log"

	"github.com/pokt-network/polylog"
)

func WithOutput(output io.Writer) polylog.LoggerOption {
	return func(logger polylog.Logger) {
		log.SetOutput(output)
	}
}

func WithLevel(level Level) polylog.LoggerOption {
	return func(logger polylog.Logger) {
		logger.(*stdLogLogger).level = level
	}
}
