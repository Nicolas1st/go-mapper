package auth

import "net/http"

type AuthResource struct {
	sessionStorage SessionStorageInterface
	database       DatabaseInterface
}

func newAuthResource(sessionStorage SessionStorageInterface, database DatabaseInterface) *AuthResource {
	return &AuthResource{
		sessionStorage: sessionStorage,
		database:       database,
	}
}

func NewAuthRouter(sessionStorage SessionStorageInterface, database DatabaseInterface) *http.ServeMux {
	authResource := newAuthResource(sessionStorage, database)
	router := http.NewServeMux()

	router.HandleFunc("/logout", authResource.logout)
	router.HandleFunc("/authenticate", authResource.authenticate)

	return router
}
