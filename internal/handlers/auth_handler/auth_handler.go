package authhandler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	usermodel "github.com/nullrish/task-manager-go/internal/models/user_model"
	authservice "github.com/nullrish/task-manager-go/internal/services/auth_service"
)

type Handler struct {
	s *authservice.Service
}

func NewHandler(s *authservice.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) RegisterUser(c fiber.Ctx) error {
	user := new(usermodel.UserRequest)
	if err := c.Bind().Body(user); err != nil {
		return err
	}
	err := h.s.RegisterUser(c, user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *Handler) LoginUser(c fiber.Ctx) error {
	user := new(usermodel.UserRequest)
	if err := c.Bind().Body(user); err != nil {
		return err
	}
	token, err := h.s.LoginUser(c, user)
	if err != nil {
		return err
	}
	return c.SendString(token)
}

func (h *Handler) RefreshToken(c fiber.Ctx) error {
	idParam := c.Params("id", "")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}
	token, err := h.s.GenerateRefreshToken(c, userID)
	if err != nil {
		return err
	}
	return c.SendString(token)
}
