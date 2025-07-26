package controller

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
  return c.SendString("YES, this is working!")
}

func Hello(c *fiber.Ctx) error {
  name := c.Params("name")
  return c.SendString("Hello " + name + "!")
}

func CreateUser(c *fiber.Ctx) error {
  type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
  }

  var user User

  if err := c.BodyParser(&user); err != nil {
    return c.Status(400).SendString("Bad Request: " + err.Error())
  }

  if user.Name == "" || user.Email == "" {
    return c.Status(400).SendString("Bad Request: Name and Email are required")
  }

  return c.Status(201).JSON(fiber.Map{
    "message": "User created successfully",
    "user": fiber.Map{
      "name":  user.Name,
      "email": user.Email,
    },
  })
}
