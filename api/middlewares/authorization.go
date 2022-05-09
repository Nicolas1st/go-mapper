package middlewares

import (
	"net/http"
	"yaroslavl-parkings/persistence/model"
)

func NewOnlyAdmin(redirectIfNot http.HandlerFunc) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var user = r.Context().Value(User).(*model.User)

			if !user.IsAdmin() {
				redirectIfNot(w, r)
				return
			}

			next(w, r)
		}
	}
}
