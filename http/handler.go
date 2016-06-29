package http

import (
	"net/http"
	"encoding/json")

func Handler() http.Handler {
	mux := http.NewServeMux()
	handler := handleHelp(mux)
	return handler
}

func respondOk(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if body == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(body)
	}
}