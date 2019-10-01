package route

import (
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

func TestTidyPath(t *testing.T) {
	path := "asd"

	if tidyPath(path) != "/asd" {
		t.Errorf("expected `/asd`, got %s", tidyPath(path))
	}
}
