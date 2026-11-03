package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"FlashHelper/internal/domain/user"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) user.Repository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	query := `
		INSERT INTO users (
			username,
			password_hash,
			role
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.PasswordHash,
		user.Role,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*user.User, error) {
	query := `
		SELECT
			id,
			username,
			password_hash,
			role
		FROM users
		WHERE id = $1
	`

	u := &user.User{}

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.Role,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, user.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	query := `
		SELECT
			id,
			username,
			password_hash,
			role
		FROM users
		WHERE username = $1
	`

	u := &user.User{}

	err := r.pool.QueryRow(ctx, query, username).Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.Role,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, user.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf(
			"get user by username: %w",
			err,
		)
	}

	return u, nil
}

func (r *UserRepository) List(ctx context.Context) ([]*user.User, error) {
	query := `
		SELECT
			id,
			username,
			password_hash,
			role
		FROM users
		ORDER BY id ASC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	defer rows.Close()

	users := make([]*user.User, 0)

	for rows.Next() {
		user := &user.User{}

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.PasswordHash,
			&user.Role,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"scan user: %w",
				err,
			)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate users: %w",
			err,
		)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET
			username = $1,
			password_hash = $2
		WHERE id = $3
	`

	result, err := r.pool.Exec(
		ctx,
		query,
		u.Username,
		u.PasswordHash,
		u.ID,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return user.ErrNotFound
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return user.ErrNotFound
	}

	log.Debug().
		Int("user_id", id).
		Msg("User deleted")

	return nil
}
