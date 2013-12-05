package web

import (
	"fmt"
	"net/http"
	"path"
)

// NewRouter returns a new router instance.
func newRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route), strictSlash: true}
}

// Router registers routes to be matched and dispatches a handler.
type Router struct {
	// Parent route, if this is a subrouter.
	parent parentRoute
	// Routes to be matched, in order.
	routes []*Route
	// Routes by name for URL building.
	namedRoutes map[string]*Route
	// See Router.StrictSlash(). This defines the flag for new routes.
	strictSlash bool
}

// Match matches registered routes against the request.
func (r *Router) Match(req *http.Request, match *RouteMatch) bool {
	for _, route := range r.routes {
		if route.Match(req, match) {
			return true
		}
	}
	return false
}

// ServeHTTP dispatches the handler registered in the matched route.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := Context{Request: req, ResponseWriter: w, Params: map[string]string{}}
	// Clean path to canonical form and redirect.
	if p := cleanPath(req.URL.Path); p != req.URL.Path {
		ctx.RedirectPermanent(p)
		return
	}

	req.ParseForm()
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			ctx.Params[k] = v[0]
		}
	}

	var match RouteMatch
	var handler Handler
	var pathVars []string
	if r.Match(req, &match) {
		handler = match.Handler
		pathVars = match.PathVars
		setVars(req, match.Vars)
		setCurrentRoute(req, match.Route)

		for k, v := range match.Vars {
			ctx.Params[k] = v
		}
	}
	if handler == nil {
		ctx.Abort(404, http.StatusText(http.StatusNotFound))
		return
	}
	routeHandler(&ctx, handler, pathVars)
}

// Get returns a route registered with the given name.
func (r *Router) Get(name string) *Route {
	return r.getNamedRoutes()[name]
}

// GetRoute returns a route registered with the given name. This method
// was renamed to Get() and remains here for backwards compatibility.
func (r *Router) GetRoute(name string) *Route {
	return r.getNamedRoutes()[name]
}

// StrictSlash defines the slash behavior for new routes.
//
// When true, if the route path is "/path/", accessing "/path" will redirect
// to the former and vice versa.
//
// Special case: when a route sets a path prefix, strict slash is
// automatically set to false for that route because the redirect behavior
// can't be determined for prefixes.
func (r *Router) StrictSlash(value bool) *Router {
	r.strictSlash = value
	return r
}

// ----------------------------------------------------------------------------
// parentRoute
// ----------------------------------------------------------------------------

// getNamedRoutes returns the map where named routes are registered.
func (r *Router) getNamedRoutes() map[string]*Route {
	if r.namedRoutes == nil {
		if r.parent != nil {
			r.namedRoutes = r.parent.getNamedRoutes()
		} else {
			r.namedRoutes = make(map[string]*Route)
		}
	}
	return r.namedRoutes
}

// getRegexpGroup returns regexp definitions from the parent route, if any.
func (r *Router) getRegexpGroup() *routeRegexpGroup {
	if r.parent != nil {
		return r.parent.getRegexpGroup()
	}
	return nil
}

// ----------------------------------------------------------------------------
// Route factories
// ----------------------------------------------------------------------------

// NewRoute registers an empty route.
func (r *Router) NewRoute() *Route {
	route := &Route{parent: r, strictSlash: r.strictSlash}
	r.routes = append(r.routes, route)
	return route
}

// Handle registers a new route with a matcher for the URL path.
// See Route.Path() and Route.Handler().
func (r *Router) Handle(path string, handler Handler) *Route {
	return r.NewRoute().Path(path).Handler(handler)
}

// Headers registers a new route with a matcher for request header values.
// See Route.Headers().
func (r *Router) Headers(pairs ...string) *Route {
	return r.NewRoute().Headers(pairs...)
}

// Host registers a new route with a matcher for the URL host.
// See Route.Host().
func (r *Router) Host(tpl string) *Route {
	return r.NewRoute().Host(tpl)
}

// MatcherFunc registers a new route with a custom matcher function.
// See Route.MatcherFunc().
func (r *Router) MatcherFunc(f MatcherFunc) *Route {
	return r.NewRoute().MatcherFunc(f)
}

// Methods registers a new route with a matcher for HTTP methods.
// See Route.Methods().
func (r *Router) Methods(methods ...string) *Route {
	return r.NewRoute().Methods(methods...)
}

// Path registers a new route with a matcher for the URL path.
// See Route.Path().
func (r *Router) Path(tpl string) *Route {
	return r.NewRoute().Path(tpl)
}

// PathPrefix registers a new route with a matcher for the URL path prefix.
// See Route.PathPrefix().
func (r *Router) PathPrefix(tpl string) *Route {
	return r.NewRoute().PathPrefix(tpl)
}

// Queries registers a new route with a matcher for URL query values.
// See Route.Queries().
func (r *Router) Queries(pairs ...string) *Route {
	return r.NewRoute().Queries(pairs...)
}

// Schemes registers a new route with a matcher for URL schemes.
// See Route.Schemes().
func (r *Router) Schemes(schemes ...string) *Route {
	return r.NewRoute().Schemes(schemes...)
}

// RouteMatch stores information about a matched route.
type RouteMatch struct {
	Route    *Route
	Handler  Handler
	PathVars []string
	Vars     map[string]string
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

// cleanPath returns the canonical path for p, eliminating . and .. elements.
// Borrowed from the net/http package.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

// uniqueVars returns an error if two slices contain duplicated strings.
func uniqueVars(s1, s2 []string) error {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				return fmt.Errorf("mux: duplicated route variable %q", v2)
			}
		}
	}
	return nil
}

// mapFromPairs converts variadic string parameters to a string map.
func mapFromPairs(pairs ...string) (map[string]string, error) {
	length := len(pairs)
	if length%2 != 0 {
		return nil, fmt.Errorf(
			"mux: number of parameters must be multiple of 2, got %v", pairs)
	}
	m := make(map[string]string, length/2)
	for i := 0; i < length; i += 2 {
		m[pairs[i]] = pairs[i+1]
	}
	return m, nil
}

// matchInArray returns true if the given string value is in the array.
func matchInArray(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// matchMap returns true if the given key/value pairs exist in a given map.
func matchMap(toCheck map[string]string, toMatch map[string][]string,
	canonicalKey bool) bool {
	for k, v := range toCheck {
		// Check if key exists.
		if canonicalKey {
			k = http.CanonicalHeaderKey(k)
		}
		if values := toMatch[k]; values == nil {
			return false
		} else if v != "" {
			// If value was defined as an empty string we only check that the
			// key exists. Otherwise we also check for equality.
			valueExists := false
			for _, value := range values {
				if v == value {
					valueExists = true
					break
				}
			}
			if !valueExists {
				return false
			}
		}
	}
	return true
}
