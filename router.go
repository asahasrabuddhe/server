package server

import "net/http"

type Handler func (ctx Context)

type Router struct {
	routes map[string]string
	http.Handler
}

func (r *Router) GET(path string, handler Handler) {
	r.handler(http.MethodGet, path, handler)
}

func (r *Router) handler(method, path string, handler Handler) {

}

func (r Router) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}

