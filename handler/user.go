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

func (t *User) User(c Context) {
	ctx := c.Request.Context()
	rsp := response.User{}
	JSON(ctx, c.ResponseWriter, rsp, http.StatusOK)
}

func (t *User) ListUser(c Context) {
	ctx := c.Request.Context()
	rsp := response.Users{}
	JSON(ctx, c.ResponseWriter, rsp, http.StatusOK)
}

func (t *User) AddUser(c Context) {
	ctx := c.Request.Context()

	user := &form.User{}
	if err := Bind(c.Request, user); err != nil {
		JSON(ctx, c.ResponseWriter, &ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	rsp := response.User{}

	JSON(ctx, c.ResponseWriter, rsp, http.StatusOK)
}
