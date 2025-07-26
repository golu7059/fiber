package routes

import (
  "github.com/gofiber/fiber/v2"
  "fiber/controller"
)

func SetUpRoutes(app *fiber.App) {
  api := app.Group("/api")

  api.Get("/", controller.Home)
  api.Get("/hello/:name", controller.Hello)
  api.Post("/users", controller.CreateUser)
}

