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
func TransformGorilla(r *mux.Router, group *Group) *mux.Router {
	var s *mux.Router

	if group.prefix != "" {
		s = r.PathPrefix(group.prefix).Subrouter()
	} else {
		s = r.NewRoute().Subrouter()
	}

	for _, mw := range group.middleware {
		s.Use(mux.MiddlewareFunc(mw))
	}

	for _, route := range group.routes {
		path := tidyPath(route.URI)
		s.Methods(route.Methods...).Path(path).Handler(route.Handler)
	}

	return s
}
