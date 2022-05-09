package middlewares

import (
	"context"
	"net/http"
	"yaroslavl-parkings/api/auth"
	"yaroslavl-parkings/persistence/model"
)

type SessionStorageInterface interface {
	IsSessionValid(sessionToken string) (*model.Session, bool)
}

type user string

var User user = "User"

func NewOnlyAuth(
	sessionStorage SessionStorageInterface,
	redirectIfNot http.HandlerFunc,
) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(auth.AuthCookieName)
			if err != nil {
				redirectIfNot(w, r)
				return
			}

			session, valid := sessionStorage.IsSessionValid(cookie.Value)
			if !valid {
				redirectIfNot(w, r)
				return
			}

			// there should be no error, since the session has already been checked
			r = r.WithContext(context.WithValue(r.Context(), User, session.User))

			next(w, r)
		}
	}
}

func NewOnlyNotAuth(redirectIfNot http.HandlerFunc) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie(auth.AuthCookieName)
			if err == nil {
				redirectIfNot(w, r)
				return
			}

			next(w, r)
		}
	}
}
