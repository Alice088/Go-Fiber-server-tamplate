package logger

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"runtime/debug"
)

func PrettyStackTraceHandler(logger *slog.Logger) func(c *fiber.Ctx, e interface{}) {
	return func(_ *fiber.Ctx, e interface{}) {
		logger.Error(fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack()))
	}
}
