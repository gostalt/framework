package route

import (
	"testing"

	"github.com/gostalt/framework/route/middleware"
)

func TestCollectionReturnsANewRouteGroup(t *testing.T) {
	h := new(TestStringer)

	g := Collection(
		Get("/", h),
	)

	if len(g.routes) != 1 {
		t.Errorf("expected 1 route, got %d", len(g.routes))
	}
}

func TestCollectionCanHaveMiddleware(t *testing.T) {
	g := Collection().Middleware(middleware.AddURIParametersToRequest)

	if len(g.middleware) != 1 {
		t.Errorf("expected 1 middleware item, got %d", len(g.middleware))
	}
}
