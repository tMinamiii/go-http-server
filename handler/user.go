package handler

import (
	"net/http"

	"github.com/tMinamiii/go-http-server/form"
	"github.com/tMinamiii/go-http-server/response"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (t *User) User(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsp := response.User{}
	JSON(ctx, w, rsp, http.StatusOK)
}

func (t *User) ListUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsp := response.Users{}
	JSON(ctx, w, rsp, http.StatusOK)
}

func (t *User) AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := &form.User{}
	if err := Bind(r, user); err != nil {
		JSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	rsp := response.User{}

	JSON(ctx, w, rsp, http.StatusOK)
}
