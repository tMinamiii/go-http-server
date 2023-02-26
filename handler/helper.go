package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func Bind(r *http.Request, f any) error {
	switch r.Method {
	case http.MethodGet, http.MethodDelete:
		tmp, err := json.Marshal(r.URL.Query())
		if err != nil {
			return err
		}
		err = json.Unmarshal(tmp, f)
		if err != nil {
			return err
		}
	case http.MethodPost, http.MethodPut:
		if err := json.NewDecoder(r.Body).Decode(f); err != nil {
			return fmt.Errorf("failed to bind")
		}
	}
	return nil
}

func JSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			fmt.Printf("write error response error: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
