package main

import (
	"net/http"

	"github.com/nekipelovaa/collectMetrics/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`POST /update/{type}/{name}/{value}`, handlers.AddMetric)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
