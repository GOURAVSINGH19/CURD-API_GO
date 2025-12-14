package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusError = "Error"
	StatusOk    = "Ok"
)

func WriteJson(w http.ResponseWriter, data interface{}, statuscode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func Validater(err validator.ValidationErrors) Response {
	var errorMsg []string

	for _, err := range err {
		switch err.ActualTag() {
		case "required":
			errorMsg = append(errorMsg, "Field is required")
		case "email":
			errorMsg = append(errorMsg, "Invalid email format")
		default:
			errorMsg = append(errorMsg, "Validation error")
		}
	}

	return Response{
		Status: StatusError,
		Error:  "Validation Error: " + errorMsg[0],
	}
}
