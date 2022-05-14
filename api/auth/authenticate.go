package auth

import (
	"errors"
	"fmt"
	"net/http"
	"yaroslavl-parkings/data/sessionstorer"
	"yaroslavl-parkings/data/user"
)

func (resource *AuthDependencies) Authenticate(w http.ResponseWriter, r *http.Request) error {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// check wheter the u with this name already exists
	u, err := resource.database.GetUserByName(username)
	if err != nil {
		fmt.Println("username")
		return err
	}

	// check user password
	if !user.IsCorrectPassword(u, password) {
		fmt.Println("password")
		return errors.New("could not authenticate")
	}

	// create new session on the server
	session := sessionstorer.NewSession(u)
	token, expiryTime := resource.sessionStorage.StoreSession(session)

	// provide the user with the session token
	SetAuthCookie(w, token, expiryTime)

	return nil
}
