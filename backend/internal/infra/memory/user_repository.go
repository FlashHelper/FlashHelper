package memory

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"FlashHelper/internal/config"
	"FlashHelper/internal/domain/user"
)

// UserRepository holds the accounts allowed to log in.
// There is no database, so the account is seeded from configuration.
type UserRepository struct {
	users map[string]*user.User
}

func NewUserRepository(cfg *config.Config) (user.Repository, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(cfg.AuthPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("hash auth password: %w", err)
	}

	admin := &user.User{
		ID:           1,
		Username:     cfg.AuthUsername,
		PasswordHash: string(hashedPassword),
		Role:         "ADMIN",
	}

	return &UserRepository{
		users: map[string]*user.User{
			admin.Username: admin,
		},
	}, nil
}

func (r *UserRepository) GetByUsername(_ context.Context, username string) (*user.User, error) {
	u, ok := r.users[username]
	if !ok {
		return nil, user.ErrNotFound
	}

	copied := *u

	return &copied, nil
}
