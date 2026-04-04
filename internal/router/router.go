package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/nullrish/task-manager-go/internal/handler"
	"github.com/nullrish/task-manager-go/internal/middleware"
	"github.com/nullrish/task-manager-go/internal/repository"
	"github.com/nullrish/task-manager-go/internal/service"
)

func ConfigureRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	r := api.Group("/auth")

	// Register dependencies for auth/user end points.
	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	authServ := service.NewAuthService(userRepo, tokenRepo)
	authHandler := handler.NewAuthHandler(authServ)

	// Register endpoints for auth.
	r.Post("/register", authHandler.RegisterUser)
	r.Post("/login", authHandler.LoginUser)
	r.Post("/token/:type/:userID", authHandler.GenerateToken)

	// Configure jwt middelware for task related routes.
	app.Use(middleware.AuthMiddleware())

	r = api.Group("/task")

	// Register dependecnies for task end points.
	taskRepo := repository.NewTaskRepository(db)
	taskServ := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskServ)

	// Register endpoints for tasks.
	r.Post("/create", taskHandler.CreateTask)
	r.Put("/update", taskHandler.UpdateTask)
	r.Get("/by-task-id/:id", taskHandler.GetTask)
	r.Get("/by-user-id/:id", taskHandler.GetUserTasks)
	r.Get("/all", taskHandler.GetTasks)
	r.Delete("/:id", taskHandler.DeleteTask)

	// Test JWT middleware restriction without Authorization Header
	r.Get("/restricted", func(c fiber.Ctx) error {
		return c.SendString("This one is restricted")
	})
}
