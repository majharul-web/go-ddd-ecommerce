package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]string{"error": message})
}
