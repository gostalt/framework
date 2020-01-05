package route

import (
	"github.com/gorilla/mux"
)

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

	// If the Collection has a prefix applied to it, then create
	// all the routes within it using the prefix.
	if group.prefix != "" {
		s = r.PathPrefix(group.prefix).Subrouter()
	} else {
		s = r.NewRoute().Subrouter()
	}

	// Gorilla allows middleware to be added to an entire subrouter,
	// so any middleware that are defined against the collection
	// are added to the subrouter here.
	for _, mw := range group.middleware {
		s.Use(mux.MiddlewareFunc(mw))
	}

	// Iterate through each route in the collection and add it
	// to the gorilla mux instance.
	for _, route := range group.routes {
		// First, ensure the path is well formed. This allows
		// users to omit the leading slash on a route.
		path := tidyPath(route.URI)

		// Then, iterate through any middleware that are defined
		// on the individual route and use each to wrap the
		// handler. This enables per-route middleware.
		handler := route.Handler
		for _, mw := range route.middleware {
			handler = mw(handler)
		}
		s.Methods(route.Methods...).Path(path).Handler(handler)
	}

	return s
}
