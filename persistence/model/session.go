package model

import "time"

const ExpiryPeriod time.Duration = 30 * time.Minute

type Session struct {
	UserID    uint
	User      *User
	ExpiresAt time.Time
}

func NewSession(user *User) *Session {
	return &Session{
		UserID: user.ID,
		User:   user,
		// keep the user logged in for 30 minutes
		ExpiresAt: time.Now().Add(ExpiryPeriod),
	}
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
