package logger

import (
	"fmt"
	"log/slog"
	"os"
	"star_trade/backend/internal/app/infrastructure/configs"
	"time"

	"github.com/lmittmann/tint"
	multiSlog "github.com/samber/slog-multi"
)

func SetupLogger() *slog.Logger {
	var log *slog.Logger
	var logFile *os.File
	var err error

	switch os.Getenv("ENV") {
	case configs.EnvProd:
		logFile, err = os.OpenFile(os.Getenv("BUILD_LOG_PATH"), os.O_RDWR|os.O_APPEND, 0644)
	default:
		logFile, err = os.OpenFile(os.Getenv("LOG_PATH"), os.O_RDWR|os.O_APPEND, 0644)
	}

	if err != nil {
		wd, _ := os.Getwd()

		fmt.Println(err.Error(), wd)
		os.Exit(1)
	}

	switch os.Getenv("ENV") {
	case configs.EnvProd:
		log = slog.New(multiSlog.Fanout(
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
	default:
		log = slog.New(multiSlog.Fanout(
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

	return log
}
