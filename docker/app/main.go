package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("served root path /"))
	})
	r.Get("/unreachable/path", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("unreachable path /unreachable"))
	})
	r.Get("/redirected/path", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "none"
		}
		w.Write([]byte("redirected path /redirected/path and name parameter is " + name))
	})
	http.ListenAndServe(":8080", r)
}
