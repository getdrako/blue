package server

import (
	"drako/pkg/globals"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	return &Server{
		app: fiber.New(),
	}
}

func (s *Server) Start() {
	defer globals.Logger.Info("server has been started")

	s.app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	s.app.Listen(":3333", fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}
