package auth

import "net/http"

func (resource *AuthDependencies) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(AuthCookieName)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return
	}

	resource.sessionStorage.RemoveSession(cookie.Value)

	RemoveAuthCookie(w)
}
