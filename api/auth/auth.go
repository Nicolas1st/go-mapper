package auth

type AuthDependencies struct {
	sessionStorage SessionStorageInterface
	database       DatabaseInterface
}

func NewAuthHandlers(sessionStorage SessionStorageInterface, database DatabaseInterface) *AuthDependencies {
	return &AuthDependencies{
		sessionStorage: sessionStorage,
		database:       database,
	}
}
