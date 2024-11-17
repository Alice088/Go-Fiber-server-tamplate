package server_builder

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/sessions"
	"log/slog"
)

type Server struct {
	Port    string
	App     *fiber.App
	Logger  *slog.Logger
	Session *sessions.FilesystemStore
}
