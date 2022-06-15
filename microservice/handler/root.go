package handler

import (
	"net/http"
	"os"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	host, _ := os.Hostname()
	w.Write(toJSON(successResponse{
		Message: http.StatusText(http.StatusOK),
		Data: map[string]string{
			"msg":  "Welcome to Microservice with Golang",
			"host": host,
		},
	}))
}
