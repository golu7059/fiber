package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var JwtSecret = []byte("This is the secret key___ for auth")

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: JwtSecret,
		ContextKey: "user",
	})
}
