package response

import (
	"encoding/json"
	"net/http"
)

func Json(data any, w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
