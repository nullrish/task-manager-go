package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/nullrish/task-manager-go/internal/handler"
	"github.com/nullrish/task-manager-go/internal/repository"
	"github.com/nullrish/task-manager-go/internal/service"
)

func configureAuthRoutes(r fiber.Router, db *sql.DB) {
	repo := repository.NewUserRepository(db)
	s := service.NewAuthService(repo)
	h := handler.NewAuthHandler(s)
	r.Post("/register", h.RegisterUser)
	r.Post("/login", h.LoginUser)
	r.Post("/refresh", h.RefreshToken)
}
