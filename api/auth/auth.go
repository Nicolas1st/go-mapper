package auth

import "net/http"

type AuthDependencies struct {
	sessionStorage SessionStorageInterface
	database       DatabaseInterface
}

type AuthHandlers struct {
	LoginHandler  func(w http.ResponseWriter, r *http.Request) error
	LogoutHandler func(w http.ResponseWriter, r *http.Request) error
}

func NewAuthHandlers(sessionStorage SessionStorageInterface, database DatabaseInterface) *AuthHandlers {
	dependencies := &AuthDependencies{
		sessionStorage: sessionStorage,
		database:       database,
	}

	return &AuthHandlers{
		LoginHandler:  dependencies.Authenticate,
		LogoutHandler: dependencies.Logout,
	}
}
