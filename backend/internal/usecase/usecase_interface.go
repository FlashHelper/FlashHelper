package usecase

import (
	"context"
	"errors"

	"FlashHelper/internal/domain/user"
)

var (
	ErrUsernameExists  = errors.New("username already exists")
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
)

type UserService interface {
	Create(ctx context.Context, u *user.User) error
	GetByID(ctx context.Context, id int) (*user.User, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
	List(ctx context.Context) ([]*user.User, error)
	Update(ctx context.Context, u *user.User) error
	Delete(ctx context.Context, id int) error
}
