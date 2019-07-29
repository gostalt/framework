package route

import "github.com/gorilla/mux"

func tidyPath(path string) string {
	if path[0] != '/' {
		path = "/" + path
	}

	return path
}

// TransformGorilla registers a Collection of routes against the
// Gorilla mux router. This enables an independent interface of
// route registration and allows the implementation to be swapped
// out later without affecting registered routes.
func TransformGorilla(r *mux.Router, routes Collection) {
	for _, route := range routes {
		path := tidyPath(route.URI)
		r.Methods(route.Methods...).Path(path).Handler(route.Handler)
	}
}
