package logger

import (
	"io"
	"log/slog"
	"os"
	"ozone/internal/utils/logger/handler/slogpretty"
)

const (
	envLocal = "local"
	endDev   = "dev"
	envProd  = "prod"
)

func Set(envLevel string, file *os.File) *slog.Logger {

	var log *slog.Logger

	var w io.Writer
	if file != nil {
		w = io.MultiWriter(os.Stdout, file)
	} else {
		w = os.Stdout
	}
	switch envLevel {
	case envLocal:
		{
			println("Loca")
			log = newPrettySlogger()

		}
	case endDev:
		{

			log = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
		}
	case envProd:
		{
			log = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
		}

	}

	return log
}

func newPrettySlogger() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
