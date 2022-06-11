package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/shandysiswandi/learn-kubernetes/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Root)
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(handler.MethodNotAllowed)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println("server error:", err)
		os.Exit(0)
	}
}
