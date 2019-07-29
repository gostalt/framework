package route

import (
	"net/http"
)

// Definition defines a new route for the application.
type Definition struct {
	Methods []string
	Handler http.Handler
	URI     string
}

// Redirect takes a `from` URI and a `to` URI and creates a new
// http.Handler to permanently redirect users to the new URI.
func Redirect(from string, to string) Definition {
	return createDefinitionFromFunc(
		from,
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, to, http.StatusPermanentRedirect)
		},
		http.MethodGet,
	)
}

// Get creates a GET route using the given URI and http.Handler.
func Get(uri string, handler http.Handler) Definition {
	return createDefinition(uri, handler, http.MethodGet)
}

// Post creates a POST route using the given URI and http.Handler.
func Post(uri string, handler http.Handler) Definition {
	return createDefinition(uri, handler, http.MethodPost)
}

func createDefinition(uri string, handler http.Handler, methods ...string) Definition {
	return Definition{
		Methods: methods,
		Handler: handler,
		URI:     uri,
	}
}

func createDefinitionFromFunc(uri string, handler http.HandlerFunc, methods ...string) Definition {
	return createDefinition(uri, http.Handler(handler), methods...)
}
