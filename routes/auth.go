package routes

import (
	"fiber/handellers"
	"fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", handellers.Login)
	api.Get("/protected", middleware.Protected(), handellers.Protected)
}
