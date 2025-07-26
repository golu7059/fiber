package main

import (
  "fmt"
  "time"

  "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)
func main() {
  app := fiber.New();

  // implementing inbuilt middlewares
  app.Use(logger.New())
  app.Use(recover.New())

  // this is my custome middleware
  app.Use(func(c *fiber.Ctx) error{
    start := time.Now()
    err := c.Next() // call the next handler
    duration := time.Since(start)
    fmt.Printf("Request %s took %v\n", c.OriginalURL(), duration)
    return err
  })
  
  // custome middleware to add in header
  app.Use(func(c *fiber.Ctx) error{
    c.Set("X-powered-By", "GFG-Fiber")
    return c.Next()
  })

  // Roures
  app.Get("/", func(c *fiber.Ctx) error{
    return c.SendString("YES ! I am working")
  })

  app.Get("/panic", func(c *fiber.Ctx) error {
    panic("Simulated crash!") // let's see middleware will work or not
  })

  app.Listen(":3000")
}
