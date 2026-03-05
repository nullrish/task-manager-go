package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nullrish/task-manager-go/internal/middleware"
	"github.com/nullrish/task-manager-go/internal/model"
	"github.com/nullrish/task-manager-go/internal/repository"
	"github.com/nullrish/task-manager-go/internal/util/hashing"
	"github.com/nullrish/task-manager-go/internal/util/validator"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) RegisterUser(ctx context.Context, user *model.UserRequest) error {
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

func (s *AuthService) LoginUser(ctx context.Context, user *model.UserRequest) (string, error) {
	// If fields are empty then return error of missing field
	if user.Username == "" && user.Email == "" {
		return "", errors.New("enter username and email")
	}

	var u *model.User
	var err error
	if user.Email != "" {
		u, err = s.repo.GetUserByEmail(ctx, user.Email)
		if err != nil {
			return "", err
		}
		if u == nil {
			return "", errors.New("invalid login or password")
		}
	} else {
		u, err = s.repo.GetUserByUsername(ctx, user.Username)
		if err != nil {
			return "", err
		}
		if u == nil {
			return "", errors.New("invalid login or password")
		}
	}
	matched := hashing.CheckHashedPassword(user.Password, u.Password)
	if matched {
		return middleware.GenerateNewUserToken(u.ID)
	} else {
		return "", errors.New("invalid login or password")
	}
}

func (s *AuthService) GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	if userID.String() == "" {
		return "", errors.New("invalid user id")
	}
	return middleware.GenerateNewUserToken(userID)
}
