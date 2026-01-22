package router

import "github.com/gofiber/fiber/v3"

func ConfigureRoutes(app *fiber.App) {
	api := app.Group("/api", func(c fiber.Ctx) error {
		return c.SendString("API end points are working perfectly fine ✅")
	})

	auth := api.Group("/auth", func(c fiber.Ctx) error {
		return c.SendString("Auth end points are working perfectly fine ✅")
	})

	auth.Post("/register", func(c fiber.Ctx) error {
		return c.SendString("Register end point is working fine ✅")
	})
	auth.Post("/login", func(c fiber.Ctx) error {
		return c.SendString("Login end point is working fine ✅")
	})
	auth.Post("/refresh", func(c fiber.Ctx) error {
		return c.SendString("Refresh token end point is working fine ✅")
	})
	auth.Post("/logout", func(c fiber.Ctx) error {
		return c.SendString("Logout end point is working perfectly fine ✅")
	})
}
