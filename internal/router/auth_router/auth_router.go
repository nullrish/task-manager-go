package authrouter

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	authhandler "github.com/nullrish/task-manager-go/internal/handlers/auth_handler"
	userrepository "github.com/nullrish/task-manager-go/internal/repositories/user_repository"
	authservice "github.com/nullrish/task-manager-go/internal/services/auth_service"
)

func ConfigureAuthRoutes(auth fiber.Router, db *sql.DB) {
	repo := userrepository.NewUserRepository(db)
	service := authservice.NewAuthService(repo)
	handler := authhandler.NewHandler(service)
	auth.Post("/register", handler.RegisterUser)
	auth.Post("/login", handler.LoginUser)
	auth.Post("/refresh", handler.RefreshToken)
}
