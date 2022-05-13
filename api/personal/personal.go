package personal

import (
	"net/http"
	"yaroslavl-parkings/persistence/model"
)

type DatabaseInterface interface {
	CreateNewUser(username, email, passwordHash string, age uint, role model.Role) (*model.User, error)
	UpdateUserAge(id, newAge uint) error
	UpdateUserEmail(id uint, newEmail string) error
}

type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*model.Session, bool)
}

type personalDependencies struct {
	db       DatabaseInterface
	sessions SessionsInterface
}

type personalDataApi struct {
	UpdateUserEmail func(w http.ResponseWriter, r *http.Request) error
	UpdateUserAge   func(w http.ResponseWriter, r *http.Request) error
	CreateAccount   func(w http.ResponseWriter, r *http.Request) error
}

func NewPersonalDataApi(db DatabaseInterface, sessions SessionsInterface) *personalDataApi {
	dependencies := &personalDependencies{
		db:       db,
		sessions: sessions,
	}

	return &personalDataApi{
		CreateAccount:   dependencies.CreateAccount,
		UpdateUserEmail: dependencies.updateUserEmail,
		UpdateUserAge:   dependencies.updateUserAge,
	}
}
