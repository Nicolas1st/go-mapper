package middlewares

import "net/http"

type ApiMiddleware func(ApiHandler) http.HandlerFunc

type ApiHandler func(w http.ResponseWriter, r *http.Request) error
