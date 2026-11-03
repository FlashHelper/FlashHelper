package implement

import (
	"context"
	"errors"
	"strings"

	"FlashHelper/internal/domain/user"
	"FlashHelper/internal/usecase"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidRole = errors.New("invalid user role")

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) usecase.UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, u *user.User) error {
	if u == nil {
		return errors.New("user is nil")
	}

	u.Role = "USER"

	if err := validateUser(u); err != nil {
		return err
	}

	_, err := s.repo.GetByUsername(ctx, u.Username)
	if err == nil {
		return usecase.ErrUsernameExists
	}
	if !errors.Is(err, user.ErrNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(u.PasswordHash),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)

	return s.repo.Create(ctx, u)
}

func (s *UserService) GetByID(ctx context.Context, id int) (*user.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, usecase.ErrInvalidUsername
	}

	return s.repo.GetByUsername(ctx, username)
}

func (s *UserService) List(ctx context.Context) ([]*user.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) Update(ctx context.Context, u *user.User) error {
	if u == nil {
		return errors.New("user is nil")
	}

	if u.ID <= 0 {
		return user.ErrNotFound
	}

	existingUser, err := s.repo.GetByID(ctx, u.ID)
	if err != nil {
		return err
	}

	u.Role = existingUser.Role

	if err := validateUser(u); err != nil {
		return err
	}

	existingUsername, err := s.repo.GetByUsername(ctx, u.Username)
	if err != nil && !errors.Is(err, user.ErrNotFound) {
		return err
	}

	if existingUsername != nil && existingUsername.ID != u.ID {
		return usecase.ErrUsernameExists
	}

	return s.repo.Update(ctx, u)
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return user.ErrNotFound
	}

	return s.repo.Delete(ctx, id)
}

func validateUser(u *user.User) error {
	if u == nil {
		return errors.New("user is nil")
	}

	u.Username = strings.TrimSpace(u.Username)

	if u.Username == "" || len(u.Username) > 100 {
		return usecase.ErrInvalidUsername
	}

	if strings.TrimSpace(u.PasswordHash) == "" {
		return usecase.ErrInvalidPassword
	}

	if u.Role == "" {
		u.Role = "USER"
	}

	if u.Role != "USER" && u.Role != "ADMIN" {
		return ErrInvalidRole
	}

	return nil
}
