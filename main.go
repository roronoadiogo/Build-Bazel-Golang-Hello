package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	server := chi.NewRouter()
	server.Use(middleware.RequestID)
	server.Use(middleware.RealIP)
	server.Use(middleware.Logger)
	server.Use(middleware.Recoverer)

	server.Use(middleware.Timeout(60 * time.Second))

	server.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World with Chi"))
	})

	http.ListenAndServe(":3333", server)

}
