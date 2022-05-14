package sessionstorer

import (
	"time"

	"github.com/google/uuid"
)

// SessionStorer - used a storage for user user sessions
// session are automatically removed once they become expired
type SessionStorer struct {
	storage                 map[string]*Session
	clearFromExpiredOncePer time.Duration
	lastPurgedAt            time.Time
}

// NewSessionStorer - creates new session storer
func NewSessionStorer(clearFromExpiredOncePer time.Duration) *SessionStorer {
	return &SessionStorer{
		storage:                 make(map[string]*Session),
		clearFromExpiredOncePer: clearFromExpiredOncePer,
		lastPurgedAt:            time.Now(),
	}
}

// purgeFromExpiredSessions - removes all expired session based,
// should not be called explicitly by the user
// only for in-library use
func (s *SessionStorer) purgeFromExpiredSessions() {
	for token, session := range s.storage {
		if session.IsExpired() {
			delete(s.storage, token)
		}
	}
}

// StoreSession - stores session in memory,
// returns session token as string, and the expiration time
func (s *SessionStorer) StoreSession(session *Session) (string, time.Time) {
	// to avoid memory leaks the session are being purged
	// It's done every expiry perdiod of one cookies elapses
	// the persiod is defined in session.go
	if time.Now().After(s.lastPurgedAt.Add(s.clearFromExpiredOncePer)) {
		s.purgeFromExpiredSessions()
	}

	defer func() {
		// the token generator function for some reason can throw a panic
		// it's an inner implementation issue(?), so it's being handled here
		recover()
	}()

	// in case the already used token is being generated,
	// it's almost impossible but can happen anyway
	var sessionToken string
	for {
		sessionToken = uuid.NewString()
		if _, alreadyExists := s.storage[sessionToken]; !alreadyExists {
			break
		}
	}

	// storing the session in memory
	s.storage[sessionToken] = session

	return sessionToken, session.ExpiresAt
}

// RemoveSession - removes session from the storage
func (storage *SessionStorer) RemoveSession(sessionToken string) {
	delete(storage.storage, sessionToken)
}

// IsSessionValid checks whether the session is valid,
// it checks if it exists and is not too old
func (storage *SessionStorer) IsSessionValid(sessionToken string) (*Session, bool) {
	session, exists := storage.storage[sessionToken]
	if !exists {
		return session, false
	}

	if session.IsExpired() {
		storage.RemoveSession(sessionToken)
		return session, false
	}

	return session, true
}
