package auth

import (
	"net/http"
	"time"
)

func (resource *AuthDependencies) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(AuthCookieName)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return
	}

	resource.sessionStorage.RemoveSession(cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:    AuthCookieName,
		Value:   "",
		Path:    CookiePath,
		Expires: time.Now(),
	})
}
