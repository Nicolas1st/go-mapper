package auth

import (
	"time"
	"yaroslavl-parkings/data/sessionstorer"
	"yaroslavl-parkings/data/user"
)

type SessionStorageInterface interface {
	StoreSession(session *sessionstorer.Session) (string, time.Time)
	RemoveSession(sessionToken string)
	IsSessionValid(sessionToken string) (*sessionstorer.Session, bool)
}

type DatabaseInterface interface {
	GetUserByName(username string) (*user.User, error)
}
