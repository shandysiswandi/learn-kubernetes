package handler

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(toJSON(successResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    "Welcome to Microservice with Golang",
	}))
}
