package auth

import (
	"net/http"
	"time"
)

func (resource *AuthResource) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(AuthCookieName)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return
	}

	// removing the session by token
	resource.sessionStorage.RemoveSession(cookie.Value)

	// removing the session in the browser
	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   "",
		Path:    CookiePath,
		Expires: time.Now(),
	})
}
