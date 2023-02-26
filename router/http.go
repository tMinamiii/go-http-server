package router

import (
	"net/http"
	"strings"

	"github.com/tMinamiii/go-http-server/handler"
)

type Path string

type Handler func(c handler.Context)

func (p Path) String() string {
	return string(p)
}

type MethodHandler struct {
	Method  string
	Handler Handler
}

type HTTPRouter struct {
	routes map[Path]MethodHandler
}

func (hr *HTTPRouter) Handler() http.Handler {
	router := http.NewServeMux()
	for path, fn := range hr.routes {
		router.HandleFunc(path.String(), func(w http.ResponseWriter, r *http.Request) {
			if r.Method == fn.Method {
				params := hr.getPathParams(r.URL.Path, path.String())
				c := handler.NewContext(w, r, params)
				fn.Handler(c)
				return
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}
	return router
}

func (hr *HTTPRouter) getPathParams(requestPath string, routePath string) map[string]string {
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

func (hr *HTTPRouter) Get(path Path, fn Handler) {
	hr.routes[path] = MethodHandler{Method: http.MethodGet, Handler: fn}
}

func (hr *HTTPRouter) Post(path Path, fn Handler) {
	hr.routes[path] = MethodHandler{Method: http.MethodPost, Handler: fn}
}

func (hr *HTTPRouter) Put(path Path, fn Handler) {
	hr.routes[path] = MethodHandler{Method: http.MethodPut, Handler: fn}
}

func (hr *HTTPRouter) Delete(path Path, fn Handler) {
	hr.routes[path] = MethodHandler{Method: http.MethodDelete, Handler: fn}
}
