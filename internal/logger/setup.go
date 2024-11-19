package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	multiSlog "github.com/samber/slog-multi"
)

func SetupLogger() *slog.Logger {
	var logger *slog.Logger
	var logFile *os.File
	var err error

	switch os.Getenv("ENV") {
	case "prod":
		logFile, err = os.OpenFile("./logs/app.log", os.O_RDWR|os.O_APPEND, 0644)
	case "test":
		logFile = nil
		err = nil
	default:
		logFile, err = os.OpenFile("../../logs/app.log", os.O_RDWR|os.O_APPEND, 0644)
	}

	if err != nil {
		wd, _ := os.Getwd()
		fmt.Println(err.Error(), wd)
		os.Exit(1)
	}

	switch os.Getenv("ENV") {
	case "prod":
		logger = slog.New(multiSlog.Fanout(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{
				Level:     slog.LevelInfo,
				AddSource: true,
			}),
			tint.NewHandler(os.Stderr, &tint.Options{
				Level:      slog.LevelInfo,
				TimeFormat: time.Kitchen,
				AddSource:  true,
			}),
		))

	case "test":
		logger = slog.New(multiSlog.Fanout(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: false,
			}),
		))

	default:
		logger = slog.New(multiSlog.Fanout(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			}),
			tint.NewHandler(os.Stderr, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
				AddSource:  true,
			}),
		))
	}

	return logger
}
