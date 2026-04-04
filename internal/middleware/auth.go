package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_KEY"))},
		Extractor:    extractors.Chain(extractors.FromCookie("uusr"), extractors.FromAuthHeader("Bearer")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c fiber.Ctx, err error) error {
	return c.JSON(map[string]any{"error": err.Error()})
}
