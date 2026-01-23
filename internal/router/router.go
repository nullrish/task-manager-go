package router

import (
	"github.com/gofiber/fiber/v3"
	authrouter "github.com/nullrish/task-manager-go/internal/router/auth_router"
)

func ConfigureRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	authrouter.ConfigureAuthRoutes(auth)
}
