package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	Validator *validator.Validate
}

func NewTask(v *validator.Validate) *Task {
	return &Task{Validator: v}
}

func (t *Task) ListTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsp := struct {
		ID int64 `json:"id"`
	}{ID: 10}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

func (t *Task) AddTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	if err := t.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	rsp := struct {
		ID int64 `json:"id"`
	}{ID: 10}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
