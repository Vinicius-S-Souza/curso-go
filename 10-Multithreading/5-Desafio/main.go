package main

import (
	"net/http"

	"exemplo.com/desafio/handles"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{cep}", handles.BuscaCepHandler)

	http.ListenAndServe(":3000", r)
}
