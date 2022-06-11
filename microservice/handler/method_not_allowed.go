package handler

import "net/http"

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(toJSON(errorResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: http.StatusText(http.StatusMethodNotAllowed),
		Details: []string{},
	}))
}
