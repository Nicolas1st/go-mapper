package middlewares

import "net/http"

func BuildRedirectOnApiCallResult(
	successURL,
	failureURL string,
) ApiMiddleware {
	return func(next ApiHandler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := next(w, r)
			if err == nil {
				http.Redirect(w, r, successURL, http.StatusSeeOther)
			} else {
				http.Redirect(w, r, failureURL, http.StatusSeeOther)
			}
		}
	}
}
