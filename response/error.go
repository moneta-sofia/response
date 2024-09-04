package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func error(msg string, code int) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  code,
	}
}

func InternalServerError(msg string, data interface{}) Response {
	return error(msg, http.StatusInternalServerError)
}
func NotFound(msg string, data interface{}) Response {
	return error(msg, http.StatusNotFound)
}
func Unauthorized(msg string, data interface{}) Response {
	return error(msg, http.StatusUnauthorized)
}
func Forbidden(msg string, data interface{}) Response {
	return error(msg, http.StatusForbidden)
}
func BadRequest(msg string, data interface{}) Response {
	return error(msg, http.StatusBadRequest)
}

func (e ErrorResponse) Error() string {
	return e.Message
}
func (e ErrorResponse) StatusCode() int {
	return e.Status
}
func (e ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}
func (e ErrorResponse) GetData() interface{} {
	return nil
}
