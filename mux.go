package track

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

//Get adds HTTP GET method and the router path for the handler
func (r *Router) Get(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodGet, h)
}

//Post adds HTTP POST method and the router path for the handler
func (r *Router) Post(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodPost, h)
}

//Delete adds HTTP DELETE method and the router path for the handler
func (r *Router) Delete(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodDelete, h)
}

//Patch adds HTTP Patch method and the router path for the handler
func (r *Router) Patch(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodPatch, h)
}

//Put adds HTTP PUT method and the router path for the handler
func (r *Router) Put(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodPut, h)
}

//Head adds HTTP HEAD method and the router path for the handler
func (r *Router) Head(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodHead, h)
}

//Options adds HTTP OPTIONS method and the router path for the handler
func (r *Router) Options(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodOptions, h)
}

//Trace adds HTTP TRACE method and the router path for the handler
func (r *Router) Trace(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodTrace, h)
}

//Connect adds HTTP CONNECT method and the router path for the handler
func (r *Router) Connect(path string, h http.HandlerFunc) {

	r.handle(path, http.MethodConnect, h)
}

func (r *Router) handle(path, method string, h http.Handler) {

	if len(path) < 1 {
		panic("router path is empty")
	}
	if path[0] != '/' {
		panic("route path must start with /")
	}

	var pathSplitted []string
	if strings.HasPrefix(path, "/") {
		pathSplitted = strings.Split(path[1:], "/")

	} else {
		pathSplitted = strings.Split(path, "/")
	}

	if len(pathSplitted) > 0 {

		for _, v := range r.middlewares {
			h = v(h.ServeHTTP)
		}

		r.add(pathSplitted, method, h)
	}

}

func exampleMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("befor")
		next.ServeHTTP(w, r)
		fmt.Println("after")
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := req.URL.Path
	if strings.HasSuffix(reqPath, "/") {
		reqPath = strings.TrimSuffix(reqPath, "/")
	}
	reqParts := strings.Split(reqPath, "/")[1:]

	router := r.search(reqParts)

	if router == nil {
		http.NotFound(w, req)
		return
	}

	// if router.Method != req.Method {
	// 	methodNotAllowedHandler(w, req)
	// 	return
	// }

	for k, v := range router.Method {
		if v == req.Method {
			break
		}
		if k == len(router.Method)-1 && v != req.Method {
			methodNotAllowedHandler(w, req)
			return
		}
	}

	// var ctx context.Context

	if len(router.Pram) > 0 {

		for k, v := range router.Pram {
			ctx := context.WithValue(req.Context(), k[1:], v)
			req = req.WithContext(ctx)
		}
	}
	router.Value.ServeHTTP(w, req)

	return
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(nil)
}
