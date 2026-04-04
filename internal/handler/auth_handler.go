package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	apperr "github.com/nullrish/task-manager-go/internal/errors"
	"github.com/nullrish/task-manager-go/internal/model"
	"github.com/nullrish/task-manager-go/internal/service"
	"github.com/nullrish/task-manager-go/internal/util"
	"github.com/nullrish/task-manager-go/internal/util/validator"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) RegisterUser(c fiber.Ctx) error {
	req := &model.UserRequest{}
	if err := c.Bind().Body(req); err != nil {
		return &apperr.ValidationError{Field: "", Message: "Invalid input"}
	}

	if req.Email == "" {
		return &apperr.ValidationError{Field: "email", Message: "please enter email"}
	}
	if req.Username == "" {
		return &apperr.ValidationError{Field: "username", Message: "please enter username"}
	}
	if req.Password == "" {
		return &apperr.ValidationError{Field: "username", Message: "please enter password"}
	}

	if !validator.ValidateUsername(req.Username) {
		return &apperr.ValidationError{
			Field:   "username",
			Message: "username can only contain letters, number and underscores (3-20 characters)",
		}
	}
	if !validator.ValidateEmail(req.Email) {
		return &apperr.ValidationError{
			Field:   "email",
			Message: "invalid email address",
		}
	}
	if !validator.ValidatePassword(req.Password) {
		return &apperr.ValidationError{
			Field:   "password",
			Message: "password must be 8-32 chars, include uppercase, lowercase, number, and special char",
		}
	}

	user, err := h.service.RegisterUser(c.Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(&model.Response{
		Message: "Successfully registered user!",
		Data:    user,
	})
}

func (h *AuthHandler) LoginUser(c fiber.Ctx) error {
	req := &model.UserRequest{}
	if err := c.Bind().Body(req); err != nil {
		return &apperr.ValidationError{Field: "", Message: "Invalid input"}
	}

	if req.Username == "" && req.Email == "" {
		return &apperr.ValidationError{Field: "auth", Message: "please enter email or username"}
	}

	login, err := h.service.LoginUser(c.Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(&model.Response{
		Message: "Login was successful!",
		Data:    login,
	})
}

//
// func (h *AuthHandler) RefreshToken(c fiber.Ctx) error {
// 	idParam := c.Params("id", "")
// 	userID, err := uuid.Parse(idParam)
// 	if err != nil {
// 		return &apperr.ValidationError{Field: "id", Message: "invalid user id"}
// 	}
// 	token, err := h.service.GenerateRefreshToken(c.Context(), userID)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(&model.Response{
// 		Message: "Successfully generated refresh token!",
// 		Data:    token,
// 	})
// }

func (h *AuthHandler) GenerateToken(c fiber.Ctx) error {
	tokenType := model.TokenType(c.Params("type", "refresh"))
	if tokenType.IsValid() {
		return &apperr.ValidationError{Field: "type", Message: "invalid token id [bearer, refresh, reset, verify]"}
	}
	userIDParam := c.Params("userID", "")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return &apperr.ValidationError{Field: "userID", Message: "Invalid user id"}
	}
	token, err := util.GenerateNewUserToken(userID, tokenType.String())
	if err != nil {
		return &apperr.InternalServerError{Message: "failed to generate token"}
	}
	return c.JSON(&model.Response{
		Message: "Successfully generated token",
		Data:    token,
	})
}
