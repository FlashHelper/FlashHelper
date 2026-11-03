package usecase

import (
	"context"
	"errors"

	"FlashHelper/internal/domain/session"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUnauthorized       = errors.New("unauthorized")
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (*session.Session, error)
	Logout(ctx context.Context, sessionID string) error
	Authenticate(ctx context.Context, sessionID string) (*session.Session, error)
}
