package router

import (
	"net/http"
	"strings"
)

type Path string

func (p Path) String() string {
	return string(p)
}

type MethodHandler struct {
	Method      string
	HandlerFunc http.HandlerFunc
}

type HTTPRouter struct {
	routes map[Path]MethodHandler
}

func (hr *HTTPRouter) Handler() http.Handler {
	router := http.NewServeMux()
	for path, fn := range hr.routes {
		router.HandleFunc(path.String(), func(w http.ResponseWriter, r *http.Request) {
			if r.Method == fn.Method {
				params := hr.getParams(r.URL.Path, path.String())
				for k, v := range params {
					r.URL.Query().Add(k, v)
				}
				fn.HandlerFunc(w, r)
				return
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}
	return router
}

func (hr *HTTPRouter) getParams(requestPath string, routePath string) map[string]string {
	params := make(map[string]string)
	requestParts := strings.Split(requestPath, "/")
	routeParts := strings.Split(routePath, "/")

	for i := 0; i < len(routeParts); i++ {
		if strings.HasPrefix(routeParts[i], ":") {
			key := strings.TrimPrefix(routeParts[i], ":")
			value := requestParts[i]
			params[key] = value
		}
	}

	return params
}

func (hr *HTTPRouter) Get(path Path, fn http.HandlerFunc) {
	hr.routes[path] = MethodHandler{Method: http.MethodGet, HandlerFunc: fn}
}

func (hr *HTTPRouter) Post(path Path, fn http.HandlerFunc) {
	hr.routes[path] = MethodHandler{Method: http.MethodPost, HandlerFunc: fn}
}

func (hr *HTTPRouter) Put(path Path, fn http.HandlerFunc) {
	hr.routes[path] = MethodHandler{Method: http.MethodPut, HandlerFunc: fn}
}

func (hr *HTTPRouter) Delete(path Path, fn http.HandlerFunc) {
	hr.routes[path] = MethodHandler{Method: http.MethodDelete, HandlerFunc: fn}
}
