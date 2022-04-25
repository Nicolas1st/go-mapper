package auth

import (
	"time"
	"yaroslavl-parkings/persistence/model"
)

type SessionStorageInterface interface {
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
}

type DatabaseInterface interface {
	GetUserByUsername(username string) (*model.User, error)
}
