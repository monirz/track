package track

import (
	"net/http"
	"strings"
)

//Handle matches http.Handler signature with a extra parameter
type Handle func(http.ResponseWriter, *http.Request, map[string]string)

//Router holds the defined routes
type Router struct {
	routes []route
}
type route struct {
	Path    string
	Method  string
	Handler Handle
}

func hello(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Write([]byte("hello"))
}

//New instantiate a new Router
func New() *Router {
	return &Router{}
}

//Post handles HTTP POST request through ServeHTTP
func (r *Router) Post(path string, h Handle) {
	r.handle(path, "POST", h)
}

//Get handles HTTP POST request through ServeHTTP
func (r *Router) Get(path string, h Handle) {
	r.handle(path, "GET", h)
}

//Patch handles HTTP POST request through ServeHTTP
func (r *Router) Patch(path string, h Handle) {
	r.handle(path, "PATCH", h)
}

//Delete handles HTTP POST request through ServeHTTP
func (r *Router) Delete(path string, h Handle) {
	r.handle(path, "DELETE", h)
}

func (r *Router) handle(path string, method string, h Handle) {
	if path[0] != '/' {
		panic("route path must start with /")
	}

	rt := route{}
	rt.Path = path
	rt.Method = method
	rt.Handler = h

	r.routes = append(r.routes, rt)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := req.URL.Path
	if strings.HasSuffix(reqPath, "/") {
		reqPath = strings.TrimSuffix(reqPath, "/")
	}
	reqParts := strings.Split(reqPath, "/")[1:]

	params := make(map[string]string)
	for i, route := range r.routes {

		routeParts := strings.Split(route.Path, "/")[1:]

		if len(routeParts) != len(reqParts) || route.Method != req.Method {
			if len(r.routes)-1 == i {
				http.NotFound(w, req)
				return
			}
			continue
		}

		isParam := false
		var paramTrimmed string

		for k, v := range routeParts {

			if strings.HasPrefix(v, ":") {
				isParam = true
				paramTrimmed = strings.TrimPrefix(v, ":")
				params[paramTrimmed] = reqParts[k]

			}

			if !isParam {
				if v != reqParts[k] {
					http.NotFound(w, req)
					return
				}
			}

		}

		route.Handler(w, req, params)
		return

	}
}
