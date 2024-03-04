package utils

import (
	"encoding/json"
	"net/http"
)

// JSON response
func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Error response
func JSONErrorResponse(w http.ResponseWriter, code int, msg string) {
	JSONResponse(w, code, map[string]string{"message": msg})
}
