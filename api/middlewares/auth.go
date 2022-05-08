package middlewares

import (
	"net/http"
	"yaroslavl-parkings/api/auth"
)

type SessionStorageInterface interface {
	IsSessionValid(sessionToken string) bool
}

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

			if !sessionStorage.IsSessionValid(cookie.Value) {
				redirectIfNot(w, r)
				return
			}

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
