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
	rt := route{}
	rt.Path = path
	rt.Method = "POST"
	rt.Handler = h

	r.routes = append(r.routes, rt)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := strings.Split(req.URL.Path, "/")[1:]
	params := make(map[string]string)
	for _, route := range r.routes {

		if route.Method != req.Method {
			http.NotFound(w, req)
			return
		}

		routePath := strings.Split(route.Path, "/")[1:]
		if len(routePath) != len(routePath) {
			http.NotFound(w, req)
			return
		}

		isParam := false

		for k, v := range routePath {
			if strings.HasPrefix(v, ":") {
				isParam = true
				paramTrimmed := strings.TrimPrefix(v, ":")
				params[paramTrimmed] = reqPath[k]

			}

			if !isParam {
				if v != reqPath[k] {
					http.NotFound(w, req)
					return
				}
			}

			route.Handler(w, req, params)
		}

	}
}
