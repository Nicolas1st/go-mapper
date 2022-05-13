package api

import (
	"net/http"
	"yaroslavl-parkings/api/auth"
	"yaroslavl-parkings/persistence/model"
)

type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*model.Session, bool)
}

// IsAuth - checks whether the user is authenticated,
// return true if yes, false otherwise
func IsAuth(sessions SessionsInterface, r *http.Request) bool {
	cookie, noCookieErr := r.Cookie(auth.AuthCookieName)
	if noCookieErr != nil {
		// if there is no cookie then the user can not be authenticated
		return false
	}

	_, valid := sessions.IsSessionValid(cookie.Value)

	return valid
}

// IsAuthAndAdmin - checks whether the user is authenticated,
// return true if yes, false otherwise
func IsAuthAndAdmin(sessions SessionsInterface, r *http.Request) bool {
	cookie, noCookieErr := r.Cookie(auth.AuthCookieName)
	if noCookieErr != nil {
		// if there is no cookie then the user can not be an admin
		return false
	}

	session, valid := sessions.IsSessionValid(cookie.Value)

	return valid && session.User.IsAdmin()
}

// GetSessionIfValid - checks whether the user is authenticated,
// returns user session, and its status
// if valid - true, false otherwise
func GetSessionIfValid(sessions SessionsInterface, r *http.Request) (*model.Session, bool) {
	cookie, noCookieErr := r.Cookie(auth.AuthCookieName)
	if noCookieErr != nil {
		// if there is no cookie then the user can not be an admin
		return &model.Session{}, false
	}

	return sessions.IsSessionValid(cookie.Value)
}
