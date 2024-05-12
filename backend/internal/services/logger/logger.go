package logger

import (
	"log/slog"
	"os"
)
var (
	opts = PrettyHandlerOptions{
        SlogOpts: slog.HandlerOptions{
            Level: slog.LevelDebug,
        },
    }
    handler = NewPrettyHandler(os.Stdout, opts)
    logger = slog.New(handler)
	Info = logger.Info
	Error = logger.Error
	Warn = logger.Warn
	Debug = logger.Debug
)
