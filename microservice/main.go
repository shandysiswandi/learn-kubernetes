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

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("New request...", r.URL.Path)
			h.ServeHTTP(w, r)
		})
	})
	r.HandleFunc("/", handler.Root)
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(handler.MethodNotAllowed)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server listen and serve on port", 8000)
	if err := srv.ListenAndServe(); err != nil {
		log.Println("server error:", err)
		os.Exit(0)
	}
}
