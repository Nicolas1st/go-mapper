package auth

import (
	"net/http"
	"yaroslavl-parkings/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (resource *AuthDependencies) Authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, err := resource.database.GetUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session := model.NewSession(user)
	token, expiryTime := resource.sessionStorage.StoreSession(session)
	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   token,
		Path:    CookiePath,
		Expires: expiryTime,
	})
}
