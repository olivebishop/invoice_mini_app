package server

import (
	"github.com/gofiber/fiber/v2"

	"invoice_app/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "invoice_app",
			AppName:      "invoice_app",
		}),

		db: database.New(),
	}

	return server
}
