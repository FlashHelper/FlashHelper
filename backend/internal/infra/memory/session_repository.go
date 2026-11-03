package memory

import (
	"context"
	"sync"
	"time"

	"FlashHelper/internal/domain/session"
)

// SessionRepository stores sessions in process memory.
// All sessions are lost when the server restarts.
type SessionRepository struct {
	mu       sync.RWMutex
	sessions map[string]*session.Session
}

func NewSessionRepository() session.Repository {
	return &SessionRepository{
		sessions: make(map[string]*session.Session),
	}
}

func (r *SessionRepository) Create(_ context.Context, s *session.Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.deleteExpiredLocked(time.Now())

	copied := *s
	r.sessions[s.ID] = &copied

	return nil
}

func (r *SessionRepository) GetByID(_ context.Context, id string) (*session.Session, error) {
	r.mu.RLock()
	s, ok := r.sessions[id]
	r.mu.RUnlock()

	if !ok {
		return nil, session.ErrNotFound
	}

	if s.IsExpired(time.Now()) {
		r.mu.Lock()
		delete(r.sessions, id)
		r.mu.Unlock()

		return nil, session.ErrNotFound
	}

	copied := *s

	return &copied, nil
}

func (r *SessionRepository) Delete(_ context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.sessions[id]; !ok {
		return session.ErrNotFound
	}

	delete(r.sessions, id)

	return nil
}

// deleteExpiredLocked removes expired sessions so the map does not grow forever.
// The caller must hold the write lock.
func (r *SessionRepository) deleteExpiredLocked(now time.Time) {
	for id, s := range r.sessions {
		if s.IsExpired(now) {
			delete(r.sessions, id)
		}
	}
}
