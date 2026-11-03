package session

import "time"

type Session struct {
	ID        string
	UserID    int
	Username  string
	Role      string
	ExpiresAt time.Time
}

func (s *Session) IsExpired(now time.Time) bool {
	return !now.Before(s.ExpiresAt)
}
