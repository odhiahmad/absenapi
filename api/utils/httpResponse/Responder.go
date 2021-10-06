package httpResponse

import "net/http"

type IResponder interface {
	Data(w http.ResponseWriter, status int, message string, data interface{})
	Error(w http.ResponseWriter, status int, error string)

}
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

type ErrorResponse struct {
	ErrorID int    `json:"errorId"`
	Message string `json:"message"`
}
