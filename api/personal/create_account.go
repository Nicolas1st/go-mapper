package personal

import (
	"errors"
	"net/http"
	"strconv"
)

func (d *personalDependencies) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	repeatPassword := r.PostFormValue("repeatPassword")
	email := r.PostFormValue("email")
	age := r.PostFormValue("age")

	if password != repeatPassword {
		return errors.New("password do no match")
	}

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return err
	}

	_, err = d.db.CreateNewUser(username, email, password, uint(ageInt))
	if err != nil {
		return errors.New("could not create new user")
	}

	return nil
}
