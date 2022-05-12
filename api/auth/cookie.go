package auth

import (
	"net/http"
	"time"
)

const AuthCookieName string = "auth_session_cookie"
const CookiePath string = "/"

func RemoveAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   "",
		Path:    CookiePath,
		Expires: time.Now(),
	})
}

func SetAuthCookie(w http.ResponseWriter, token string, expiryTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   token,
		Path:    CookiePath,
		Expires: expiryTime,
	})
}
