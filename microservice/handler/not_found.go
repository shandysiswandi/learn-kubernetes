package handler

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(toJSON(errorResponse{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
		Details: []string{},
	}))
}
