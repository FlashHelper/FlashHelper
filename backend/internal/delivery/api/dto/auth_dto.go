package dto

import (
	"time"

	"FlashHelper/internal/domain/session"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,max=100" example:"admin" extensions:"x-order=1"`
	Password string `json:"password" binding:"required,max=72" example:"flashhelper-dev-password" extensions:"x-order=2"`
}

type LoginResponse struct {
	SessionID string    `json:"session_id" example:"HZd-KGRl6mSP3XTLBccL9E4EiGqdwyyngtf7-FVrk0k" extensions:"x-order=1"`
	ExpiresAt time.Time `json:"expires_at" example:"2026-11-04T21:09:38+07:00" extensions:"x-order=2"`
	Username  string    `json:"username" example:"admin" extensions:"x-order=3"`
	Role      string    `json:"role" example:"ADMIN" extensions:"x-order=4"`
}

type LoginSuccessResponse struct {
	Success bool           `json:"success" example:"true" extensions:"x-order=1"`
	Data    *LoginResponse `json:"data" extensions:"x-order=2"`
}

func ToLoginResponse(s *session.Session) *LoginResponse {
	if s == nil {
		return nil
	}

	return &LoginResponse{
		SessionID: s.ID,
		ExpiresAt: s.ExpiresAt,
		Username:  s.Username,
		Role:      s.Role,
	}
}

func NewLoginSuccessResponse(s *session.Session) LoginSuccessResponse {
	return LoginSuccessResponse{
		Success: true,
		Data:    ToLoginResponse(s),
	}
}
