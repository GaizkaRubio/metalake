package http

import (
	"net/http"
)

func handleHelp(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		respondOk(w, "navokob")
		h.ServeHTTP(w, req)
		return
	})
}