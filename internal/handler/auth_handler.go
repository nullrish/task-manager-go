package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/nullrish/task-manager-go/internal/model"
	"github.com/nullrish/task-manager-go/internal/service"
)

type AuthHandler struct {
	s *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) RegisterUser(c fiber.Ctx) error {
	user := new(model.UserRequest)
	if err := c.Bind().Body(user); err != nil {
		return err
	}
	err := h.s.RegisterUser(c, user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *AuthHandler) LoginUser(c fiber.Ctx) error {
	user := new(model.UserRequest)
	if err := c.Bind().Body(user); err != nil {
		return err
	}
	token, err := h.s.LoginUser(c, user)
	if err != nil {
		return err
	}
	return c.SendString(token)
}

func (h *AuthHandler) RefreshToken(c fiber.Ctx) error {
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
