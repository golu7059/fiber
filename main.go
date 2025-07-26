package main

import "github.com/gofiber/fiber/v2"

func main() {
  app := fiber.New();

  app.Get("/", func(c *fiber.Ctx) error{
    return c.SendString("Yes This application is working")
  })


  app.Get("/hello/:name", func(c *fiber.Ctx) error{
    name := c.Params("name")
    return c.SendString("Hello" + name)
  })

  app.Post("/user", func(c *fiber.Ctx) error{
    type User struct {
      Name string `json:"name"`
      Email string `json:"email"`
    }

    var user User
    if err := c.BodyParser(&user); err != nil {
      return c.Status(400).SendString("Invalid input")
    }

    return c.JSON(fiber.Map {
      "message": "User created", 
      "user": user,
    })
  })

  app.Listen(":3000")
}
