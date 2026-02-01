package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	authrouter "github.com/nullrish/task-manager-go/internal/router/auth_router"
)

func ConfigureRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	authrouter.ConfigureAuthRoutes(auth, db)
}
