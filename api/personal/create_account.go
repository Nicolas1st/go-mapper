package personal

import (
	"errors"
	"net/http"
	"strconv"
	"yaroslavl-parkings/api/middlewares"
	"yaroslavl-parkings/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (d *personalDependencies) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	repeatPassword := r.PostFormValue("repeatPassword")
	email := r.PostFormValue("email")
	age := r.PostFormValue("age")

	flashMessages, _ := middlewares.GetFlashMessages(r)

	if password != repeatPassword {
		flashMessages.AddFlashMessage("the passwords do not match", model.ErrorMessage)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("could not hash the password")
	}

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return err
	}

	_, err = d.db.CreateNewUser(username, email, string(passwordHash), uint(ageInt), model.Customer)
	if err != nil {
		flashMessages.AddFlashMessage(err.Error(), model.ErrorMessage)
		return errors.New("could not create new user")
	}

	return nil
}
