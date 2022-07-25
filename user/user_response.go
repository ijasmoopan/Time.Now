package user

import (
	"encoding/json"
	"net/http"
)

// ResponseOK function for giving response to a request
func ResponseOK (w http.ResponseWriter, response map[string]interface{}){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

// ResponseBR function for giving response to a request
func ResponseBR (w http.ResponseWriter, response map[string]interface{}){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&response)
}