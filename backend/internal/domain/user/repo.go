package user

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}