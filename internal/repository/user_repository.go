package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nullrish/task-manager-go/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *model.UserRequest) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, req *model.UserRequest) (*model.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, req *model.UserRequest) (*model.User, error) {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, username, email, password, created_at, updated_at;
	`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Printf("(user_repository) - [CreateUser] failed to create user %s: %v", req.Username, err)
		return nil, fmt.Errorf("CreateUser: %w", err)
	}
	return &user, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at
		FROM users WHERE username = $1
	`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Printf("(user_repository) - [GetUserByUsername] failed to get username %s: %v", username, err)
		return nil, fmt.Errorf("GetUserByUsername: %w", err)
	}
	return &user, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at
		FROM users WHERE email = $1
	`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Printf("(user_repository) - [GetUserByEmail] failed to get email %s: %v", email, err)
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, userID uuid.UUID, req *model.UserRequest) (*model.User, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, password = $3
		WHERE id = $4
		RETURNING id, username, email, password, created_at, updated_at;
	`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, req.Username, req.Email, req.Password, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Printf("(user_repository) - [UpdateUser] failed to update user %s: %v", userID, err)
		return nil, fmt.Errorf("UpdateUser: %w", err)
	}
	return &user, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("(user_repository) - [DeleteUser] failed to delete user %s: %v", userID, err)
		return fmt.Errorf("DeleteUser: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteUser - rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
