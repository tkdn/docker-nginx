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
		w.Write([]byte("GO-APP: served root path /"))
	})
	r.Get("/api/v1/path", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GO-APP: /api/v1/path"))
	})
	r.Get("/api/v1/content/path", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GO-APP: unreachable path /api/v1/content/path"))
	})
	r.Route("/api/v1/method/put", func(r chi.Router) {
		r.Put("/{putID}", func(w http.ResponseWriter, r *http.Request) {
			putID := chi.URLParam(r, "putID")
			w.Write([]byte("GO-APP: unreachable path /api/v1/content/path/put/:id + id = " + putID))
		})
	})
	http.ListenAndServe(":8080", r)
}
