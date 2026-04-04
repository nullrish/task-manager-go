package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenType string

const (
	Bearer  TokenType = "bearer"
	Refresh TokenType = "refresh"
	Reset   TokenType = "reset"
	Verify  TokenType = "verify"
)

func (t TokenType) IsValid() bool {
	switch t {
	case Bearer, Refresh, Reset, Verify:
		return true
	default:
		return false
	}
}

func (t TokenType) GetExpiryTime() time.Time {
	switch t {
	case Bearer:
		return time.Now().Add(time.Minute * 15)
	case Reset:
		return time.Now().Add(30 * time.Minute)
	case Verify:
		return time.Now().Add(5 * time.Hour)
	default:
		return time.Now().Add(time.Hour * 24 * 7)
	}
}

func (t TokenType) String() string {
	return string(t)
}

type UserToken struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	Token     string    `db:"token"`
	TokenType TokenType `db:"token_type"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
	Revoked   bool      `db:"revoked"`
}
