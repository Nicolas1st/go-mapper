package auth

import (
	"net/http"
	"yaroslavl-parkings/persistence/model"
)

func (resource *AuthResource) authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// check if user exists
	user, err := resource.database.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// for now there is no password hashing
	if password != user.PasswordHash {
		w.WriteHeader(http.StatusBadRequest)
	}

	// create session
	session := model.NewSession(user.ID)

	// store session in memory
	token, expiryTime := resource.sessionStorage.StoreSession(session)

	// set session cookie in the user's browser
	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   token,
		Path:    CookiePath,
		Expires: expiryTime,
	})
}
