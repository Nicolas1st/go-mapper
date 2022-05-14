package sessionstorer

type SessionInterface interface {
	GetUser() UserInterface
	IsExpired() bool
}

type UserInterface interface {
	IsUserAdmin() bool
}
