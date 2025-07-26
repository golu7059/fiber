package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
  "github.com/gofiber/fiber/v2/middleware/recover"

  "fiber/routes"
)

func main() {
  app := fiber.New()

  app.Use(logger.New())
  app.Use(recover.New())

  routes.SetUpRoutes(app)

  app.Listen(":3000")
}

