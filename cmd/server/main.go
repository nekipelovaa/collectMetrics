package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nekipelovaa/collectMetrics/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", handlers.AddMetric)
	r.Get("/value/{type}/{name}", handlers.GetMetric)
	r.Get("/", handlers.GetAllMetrics)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
