package router

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tMinamiii/go-http-server/handler"
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

func (r *HTTPRouter) NewHandler() http.Handler {
	router := http.NewServeMux()
	for path, fn := range r.routes {
		router.HandleFunc(path.String(), func(w http.ResponseWriter, r *http.Request) {
			if r.Method == fn.Method {
				fn.HandlerFunc(w, r)
				return
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}
	return router
}

func (r *HTTPRouter) Get(path Path, fn http.HandlerFunc) {
	r.routes[path] = MethodHandler{Method: http.MethodGet, HandlerFunc: fn}
}

func (r *HTTPRouter) Post(path Path, fn http.HandlerFunc) {
	r.routes[path] = MethodHandler{Method: http.MethodPost, HandlerFunc: fn}
}

func Router() http.Handler {
	r := &HTTPRouter{}
	ok := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}

	v := validator.New()
	task := handler.NewTask(v)

	r.Get("/health", ok)
	r.Get("/tasks", task.ListTask)
	r.Post("/tasks", task.AddTask)
	return nil
}
