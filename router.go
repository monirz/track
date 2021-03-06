package track

import (
	"net/http"
)

type Router struct {
	Children    map[string]*Router
	Value       http.Handler
	isEnd       bool
	isParam     bool
	Pram        map[string]string
	Method      []string
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

//New creates new instance
func New() *Router {

	m := make(map[string]*Router)
	params := make(map[string]string)
	return &Router{Children: m, Pram: params}
}

func (r *Router) add(vals []string, method string, h http.Handler) {

	if len(vals) < 1 {
		r.isEnd = true
		r.Value = h
		r.Method = append(r.Method, method)
		return
	}

	if _, ok := r.Children[vals[0]]; !ok {
		r.Children[vals[0]] = New()

		if vals[0][0] == ':' {
			r.Children[vals[0]].isParam = true
		}
	}

	r.Children[vals[0]].add(vals[1:], method, h)

}

func (r *Router) search(vals []string) *Router {

	if r == nil {
		return nil
	}

	curr := r

	for _, v := range vals {

		if _, ok := curr.Children[v]; !ok {

			if len(curr.Children) > 0 {

				for k := range curr.Children {

					if curr.Children[k].isParam {
						curr.Pram[k] = v
						v = k
						break
					}
				}

			} else {
				return nil

			}

		}

		curr = curr.Children[v]

		if curr == nil {
			return nil
		}
	}

	return curr
}
