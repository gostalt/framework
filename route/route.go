package route

import (
	"fmt"
	"net/http"
)

// Definition defines a new route for the application.
type Definition struct {
	Methods    []string
	Handler    http.Handler
	URI        string
	middleware []Middleware
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

// Get creates a GET route using the given URI.
func Get(uri string, handler interface{}) Definition {
	return createDefinition(uri, handler, http.MethodGet)
}

// Post creates a POST route using the given URI.
func Post(uri string, handler interface{}) Definition {
	return createDefinition(uri, handler, http.MethodPost)
}

// Put creates a PUT route using the given URI.
func Put(uri string, handler interface{}) Definition {
	return createDefinition(uri, handler, http.MethodPut)
}

// Patch creates a PATCH route using the given URI.
func Patch(uri string, handler interface{}) Definition {
	return createDefinition(uri, handler, http.MethodPatch)
}

// Delete creates a DELETE route using the given URI.
func Delete(uri string, handler interface{}) Definition {
	return createDefinition(uri, handler, http.MethodDelete)
}

// Options creates a OPTIONS route using the given URI.
func Options(uri string, handler interface{}) Definition {
	return createDefinition(uri, handler, http.MethodOptions)
}

func (d Definition) Middleware(middleware ...Middleware) Definition {
	d.middleware = middleware
	return d
}

func createDefinition(uri string, handler interface{}, methods ...string) Definition {
	var h http.Handler

	switch handler.(type) {
	case http.Handler:
		h = handler.(http.Handler)
	case fmt.Stringer:
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(handler.(fmt.Stringer).String()))
		})
	default:
		panic("cannot create definition")
	}

	return Definition{
		Methods: methods,
		Handler: h,
		URI:     uri,
	}
}

func createDefinitionFromFunc(uri string, handler http.HandlerFunc, methods ...string) Definition {
	return createDefinition(uri, http.Handler(handler), methods...)
}
