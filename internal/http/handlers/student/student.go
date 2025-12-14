package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gouravsingh19/CURD-API/internal/types"
	"github.com/gouravsingh19/CURD-API/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Student handler invoked")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, response.GeneralError(err), http.StatusBadRequest)
			return
		}

		slog.Info("Student data received")
		if err := validator.New().Struct(student); err != nil {
			validatorErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, response.Validater(validatorErrs), http.StatusBadRequest)
			return
		}
		w.Write([]byte("Student Handler"))
	}
}
