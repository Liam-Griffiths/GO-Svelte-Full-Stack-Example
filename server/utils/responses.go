package responses

import (
	"encoding/json"
	"net/http"
)

// return error message
func Error(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, map[string]string{"message": msg})
}

func Text(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, map[string]string{"message": msg})
}

func JSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
