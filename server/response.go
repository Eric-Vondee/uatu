package server

import (
	"net/http"

	"github.com/go-chi/render"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message" validate:"required"`
	Data    interface{} `json:"data,omitempty"`
}

func (a APIResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, a.Code)
	return nil
}

type APIError struct {
	APIResponse
}

func newAPIResponse(code int, message string, data interface{}) APIResponse {
	return APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
