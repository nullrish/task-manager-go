package services

import (
	"context"
	"errors"

	models "github.com/nullrish/task-manager-go/internal/models/user_model"
	repo "github.com/nullrish/task-manager-go/internal/repositories/user_repository"
	"github.com/nullrish/task-manager-go/internal/util/hashing"
	"github.com/nullrish/task-manager-go/internal/util/validator"
)

type AuthService struct {
	repo repo.UserRepository
}

func NewAuthService(repo repo.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) RegisterUser(ctx context.Context, user *models.UserRequest) error {
	// If fields are empty then return error of missing field
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing fields required")
	}
	// Validate email address
	if !validator.ValidateEmail(user.Email) {
		return errors.New("invalid email address")
	}
	// Validate username
	if !validator.ValidateUsername(user.Username) {
		return errors.New("username can only contain letters, number and underscores (3-20 characters)")
	}
	// Validate password
	if !validator.ValidatePassword(user.Password) {
		return errors.New("password must be 8-32 chars, include uppercase, lowercase, number, and special char")
	}
	// Check if username already exists.
	if existing, err := s.repo.GetUserByUsername(ctx, user.Username); err != nil {
		return err
	} else if existing != nil {
		return errors.New("username already exists")
	}
	// Check if email is already registered.
	if existing, err := s.repo.GetUserByEmail(ctx, user.Email); err != nil {
		return err
	} else if existing != nil {
		return errors.New("email already exists")
	}
	var err error
	user.Password, err = hashing.HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash the password")
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *AuthService) LoginUser(ctx context.Context, user *models.UserRequest) (string, error) {
}
