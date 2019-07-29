package route

// Group is a collection of route Definitions.
type Group struct {
	routes     []Definition
	prefix     string
	middleware []Middleware
}

// Collection creates a Group of route Definitions.
func Collection(routes ...Definition) *Group {
	return &Group{
		routes: routes,
	}
}

// Prefix adds a URI prefix to the list of Route definitions.
// For example, with a prefix of "api", "users" would become
// "/api/users".
func (g *Group) Prefix(name string) *Group {
	if name[0] != '/' {
		name = "/" + name
	}

	g.prefix = name
	return g
}

// Middleware adds Middleware on to the Group.
func (g *Group) Middleware(middleware ...Middleware) *Group {
	g.middleware = middleware
	return g
}
