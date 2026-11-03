package implement

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"FlashHelper/internal/config"
	"FlashHelper/internal/domain/session"
	"FlashHelper/internal/domain/user"
	"FlashHelper/internal/usecase"
)

const sessionIDBytes = 32

// dummyPasswordHash is compared when the username does not exist,
// so a login attempt takes the same time whether or not the user exists.
var dummyPasswordHash, _ = bcrypt.GenerateFromPassword(
	[]byte("flashhelper-dummy-password"),
	bcrypt.DefaultCost,
)

type AuthService struct {
	userRepo    user.Repository
	sessionRepo session.Repository
	sessionTTL  time.Duration
}

func NewAuthService(
	cfg *config.Config,
	userRepo user.Repository,
	sessionRepo session.Repository,
) usecase.AuthService {
	return &AuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		sessionTTL:  cfg.SessionTTL,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*session.Session, error) {
	username = strings.TrimSpace(username)

	u, err := s.userRepo.GetByUsername(ctx, username)
	if errors.Is(err, user.ErrNotFound) {
		_ = bcrypt.CompareHashAndPassword(dummyPasswordHash, []byte(password))
		return nil, usecase.ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return nil, usecase.ErrInvalidCredentials
	}

	sessionID, err := newSessionID()
	if err != nil {
		return nil, err
	}

	sess := &session.Session{
		ID:        sessionID,
		UserID:    u.ID,
		Username:  u.Username,
		Role:      u.Role,
		ExpiresAt: time.Now().Add(s.sessionTTL),
	}

	if err := s.sessionRepo.Create(ctx, sess); err != nil {
		return nil, err
	}

	return sess, nil
}

func (s *AuthService) Logout(ctx context.Context, sessionID string) error {
	err := s.sessionRepo.Delete(ctx, sessionID)
	if errors.Is(err, session.ErrNotFound) {
		return usecase.ErrUnauthorized
	}

	return err
}

func (s *AuthService) Authenticate(ctx context.Context, sessionID string) (*session.Session, error) {
	if sessionID == "" {
		return nil, usecase.ErrUnauthorized
	}

	sess, err := s.sessionRepo.GetByID(ctx, sessionID)
	if errors.Is(err, session.ErrNotFound) {
		return nil, usecase.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func newSessionID() (string, error) {
	b := make([]byte, sessionIDBytes)

	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate session id: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
