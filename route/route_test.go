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

type TestStringer int

func (TestStringer) String() string {
	return "Just Testing"
}
