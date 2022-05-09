package session

import (
	"time"
	"yaroslavl-parkings/persistence/model"

	"github.com/google/uuid"
)

type SessionStorage struct {
	sessions     map[string]*model.Session
	lastPurgedAt time.Time
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		sessions:     make(map[string]*model.Session),
		lastPurgedAt: time.Now(),
	}
}

func (storage *SessionStorage) purgeFromExpiredSessions() {
	for token, session := range storage.sessions {
		if session.IsExpired() {
			delete(storage.sessions, token)
		}
	}
}

func (storage *SessionStorage) StoreSession(session *model.Session) (string, time.Time) {
	// to avoid memory leaks the session are being purged
	// It's done every expiry perdiod of one cookies elapses
	// the persiod is defined in session.go
	if time.Now().After(storage.lastPurgedAt.Add(model.ExpiryPeriod)) {
		storage.purgeFromExpiredSessions()
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
		if _, alreadyExists := storage.sessions[sessionToken]; !alreadyExists {
			break
		}
	}

	// storing the session in memory
	storage.sessions[sessionToken] = session

	return sessionToken, session.ExpiresAt
}

func (storage *SessionStorage) RemoveSession(sessionToken string) {
	delete(storage.sessions, sessionToken)
}

// IsSessionValid checks whether the session is valid,
// it checks if it exists and is not too old
func (storage *SessionStorage) IsSessionValid(sessionToken string) (*model.Session, bool) {
	session, exists := storage.sessions[sessionToken]
	if !exists {
		return session, false
	}

	if session.IsExpired() {
		storage.RemoveSession(sessionToken)
		return session, false
	}

	return session, true
}
