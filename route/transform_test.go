package route

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestTransformGorillaAddsRoutes(t *testing.T) {
	r := mux.NewRouter()
	h := new(TestStringer)
	g := Collection(
		Get("/", h),
	)

	result := TransformGorilla(r, g)

	result.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if p, _ := route.GetPathTemplate(); p != "/" {
			t.Error("Path `/` was not added to Router")
		}

		return nil
	})
}

func TestTransformGorillaUsesGroupPrefix(t *testing.T) {
	r := mux.NewRouter()
	h := new(TestStringer)
	g := Collection(
		Get("hello", h),
	).Prefix("test")

	result := TransformGorilla(r, g)

	result.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if p, _ := route.GetPathTemplate(); p != "/test/hello" {
			t.Error("Prefixed route `/test/hello` was not added to Router")
		}

		return nil
	})
}

func TestMiddlewareChainIsAppliedToIndividualRoute(t *testing.T) {
	r := mux.NewRouter()
	h := new(TestStringer)
	g := Collection(
		Get("/test", h).Middleware(MiddlewareTester, MiddlewareTesterTwo),
	)

	result := TransformGorilla(r, g)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	result.ServeHTTP(w, req)

	if w.Body.String() != "THIS IS A TEST Just Testing" {
		t.Error("Expect middleware to be added to the route")
	}
}

func TestTidyPath(t *testing.T) {
	path := "asd"

	if tidyPath(path) != "/asd" {
		t.Errorf("expected `/asd`, got %s", tidyPath(path))
	}
}
