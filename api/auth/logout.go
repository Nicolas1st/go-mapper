package auth

import (
	"errors"
	"net/http"
)

func (resource *AuthDependencies) Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(AuthCookieName)
	if err == http.ErrNoCookie {
		return errors.New("logged out user trying to log out")
	}

	resource.sessionStorage.RemoveSession(cookie.Value)
	RemoveAuthCookie(w)

	return nil
}
