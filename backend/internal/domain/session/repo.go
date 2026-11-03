package session

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("session not found")

type Repository interface {
	Create(ctx context.Context, s *Session) error
	GetByID(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
}
