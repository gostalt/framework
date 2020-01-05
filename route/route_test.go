package route

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var testHandler = http.Handler(testHandlerFunc)

var testHandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
})

func TestCreateDefinition(t *testing.T) {
	r := createDefinition("/", testHandler)

	if _, ok := r.Handler.(http.Handler); !ok {
		t.Error("`createDefinition` didn't create a handler")
	}
}

func TestCreateDefinitionAllowsHandlerFunc(t *testing.T) {
	r := createDefinition("/", testHandlerFunc)

	if _, ok := r.Handler.(http.Handler); !ok {
		t.Error("`createDefinition` didn't create a handler from HandlerFunc")
	}
}

func TestCreateDefinitionAllowsStringer(t *testing.T) {
	s := new(TestStringer)
	r := createDefinition("/", s)
	w := httptest.NewRecorder()

	r.Handler.ServeHTTP(w, nil)

	if w.Body.String() != s.String() {
		t.Error("`createDefinition`'s return wasn't the same as the Stringer")
	}
}

func TestRedirect(t *testing.T) {
	redirect := Redirect("example", "test")

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/example", nil)

	redirect.Handler.ServeHTTP(w, r)

	if w.Code != http.StatusPermanentRedirect {
		t.Error("Expected redirect to return 308")
	}
}

func TestInvalidRouteHandlerPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic")
		}
	}()

	_ = createDefinition("/", 42)
}

func TestHTTPRouteMethodsReturnDefinition(t *testing.T) {
	methods := []func(string, interface{}) Definition{
		Get,
		Post,
		Put,
		Patch,
		Delete,
		Options,
	}

	s := new(TestStringer)

	for _, m := range methods {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Creating definition panicked")
			}
		}()

		m("/", s)
	}
}

type TestStringer int

func (TestStringer) String() string {
	return "Just Testing"
}

func MiddlewareTester(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("A TEST "))
		next.ServeHTTP(w, r)
	})
}

func MiddlewareTesterTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("THIS IS "))
		next.ServeHTTP(w, r)
	})
}
