package track

import (
	"log"
	"net/http"
	"strings"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

type middleware interface {
	Middleware(handler http.Handler) http.Handler
}

func (mw MiddlewareFunc) Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return mw(handler)
}

func (r *Router) Use(mwf ...MiddlewareFunc) {
	for _, fn := range mwf {
		r.middlewares = append(r.middlewares, fn)
	}
}

func CORSMethodMiddleware(r *Router) MiddlewareFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			methods := getMethodsForRoute(r, req).Method

			log.Println("debug mux", r.Method)

			for _, v := range methods {
				if v == http.MethodOptions {
					w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				}
			}
			next.ServeHTTP(w, req)

		})
	}
}

func getMethodsForRoute(r *Router, req *http.Request) *Router {

	vals := strings.Split(req.URL.Path[1:], "/")
	return r.search(vals)

}
