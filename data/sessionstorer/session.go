package sessionstorer

import (
	"time"
	"yaroslavl-parkings/data/user"
)

const ExpiryPeriod time.Duration = 5 * time.Minute

type Session struct {
	UserID    uint
	User      *user.User
	ExpiresAt time.Time
}

func NewSession(user *user.User) *Session {
	return &Session{
		UserID: user.ID,
		User:   user,
		// keep the user logged in for 5 minutes
		ExpiresAt: time.Now().Add(ExpiryPeriod),
	}
}

func (s *Session) GetUser() *user.User {
	return s.User
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
