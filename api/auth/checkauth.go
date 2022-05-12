package auth

import (
	"net/http"
	"yaroslavl-parkings/persistence/model"
)

func (resource *AuthDependencies) GetSessionIfLoggedIn(w http.ResponseWriter, r *http.Request) (*model.Session, bool) {
	cookie, noCookieErr := r.Cookie(AuthCookieName)
	if noCookieErr != nil {
		return &model.Session{}, false
	}

	token := cookie.Value
	session, valid := resource.sessionStorage.IsSessionValid(token)
	if !valid {
		// remove the cookie from user's browser
		RemoveAuthCookie(w)
		// remove the session from the storage
		resource.sessionStorage.RemoveSession(token)
		return &model.Session{}, false
	}

	return session, true
}
