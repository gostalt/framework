package route

import "net/http"

// Middleware must accept an http.Handler and return an http.Handler.
// More complex middleware can be a struct, with a `Handle` method
// that accepts an http.Handler and returns an http.Handler.
type Middleware func(http.Handler) http.Handler
