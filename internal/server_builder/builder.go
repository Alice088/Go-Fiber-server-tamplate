package server_builder

import (
	"RuRu/internal/api/controller"
	customMiddleware "RuRu/internal/api/middleware"
	supportSession "RuRu/internal/api/session"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	"RuRu/internal/logger"
	_ "errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"os"
)

type IBuilder interface {
	InitFiber()
	InitLogger()
	InitSession()
	InitMiddleware()
	InitRoutes()
}

type Builder struct {
	Server *Server
}

func (b Builder) InitFiber() {
	b.Server.App = fiber.New(fiber.Config{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Fiber",
		BodyLimit:         300 * 1024 * 1024,
		EnablePrintRoutes: true,
		AppName:           fmt.Sprintf("RuRu-audios v%s", os.Getenv(`API_VERSION`)),
	})
}

func (b Builder) InitLogger() {
	b.Server.Logger = logger.SetupLogger()
}

func (b Builder) InitSession() {
	b.Server.Session = supportSession.New()
}

func (b Builder) InitMiddleware() {
	b.Server.App.Get("/metrics", monitor.New())
	b.Server.App.Use(healthcheck.New())
	b.Server.App.Use(helmet.New())
	b.Server.App.Use(idempotency.New())
	b.Server.App.Use(requestid.New())
	b.Server.App.Use(pprof.New())
	b.Server.App.Use(limiter.New(limiter.Config{
		Max: 20,
	}))
	b.Server.App.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: logger.PrettyStackTraceHandler(b.Server.Logger),
	}))
	b.Server.App.Use(customMiddleware.Hmac(b.Server.Logger, b.Server.Session))
}

func (b Builder) InitRoutes() {
	v1 := b.Server.App.Group("/api").Group("/v1")

	v1Audio := v1.Group("/audio")

	v1Audio.Post("/transcription", controller.Transcription(b.Server.Logger))
}
