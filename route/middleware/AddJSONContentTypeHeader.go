package middleware

import "net/http"

// AddJSONContentTypeHeader adds a `Content-Type` header to the
// response with a value of `application/json`, meaning that
// all routes that implement this middleware return JSON.
func AddJSONContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}
