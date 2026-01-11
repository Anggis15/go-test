package halper

import (
	"encoding/json"
	"go-backend-univ/model"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(model.ErrorGeneral{
		Status:  status,
		Message: message,
	})
}
