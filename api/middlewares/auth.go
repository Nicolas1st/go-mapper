package middlewares

import (
	"net/http"
	"yaroslavl-parkings/api/auth"
)

type AuthMiddleware func(http.HandlerFunc) http.HandlerFunc

type SessionStorageInterface interface {
	IsSessionValid(sessionToken string) bool
}

func NewAuthMiddleware(sessionStorage SessionStorageInterface, redirectUnauthenticated http.HandlerFunc) AuthMiddleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(auth.AuthCookieName)
			if err == http.ErrNoCookie {
				redirectUnauthenticated(w, r)
				return
			}

			// checking by session token
			if !sessionStorage.IsSessionValid(cookie.Value) {
				redirectUnauthenticated(w, r)
				return
			}

			next(w, r)
		}
	}
}
