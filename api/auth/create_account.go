package auth

import (
	"net/http"
	"yaroslavl-parkings/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (resource *AuthDependencies) CreateAccount(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	repeatPassword := r.PostFormValue("repeatPassword")

	if password != repeatPassword {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = resource.database.CreateNewUser(username, string(passwordHash), model.Customer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
