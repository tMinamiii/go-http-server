package router

import (
	"net/http"

	"github.com/tMinamiii/go-http-server/handler"
)

func Router() http.Handler {
	r := &HTTPRouter{}
	ok := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}

	user := handler.NewUser()

	r.Get("/health", ok)
	r.Get("/users", user.ListUser)
	r.Post("/users", user.AddUser)
	r.Get("/users/:id", user.User)

	return r.Handler()
}
