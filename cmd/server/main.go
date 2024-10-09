package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nekipelovaa/collectMetrics/internal/handlers"
)

func main() {
	addr := flag.String("a", "localhost:8080", "адрес HTTP сервера")
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("Неизвестный флаг:", flag.Args())
		return
	}

	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", handlers.AddMetric)
	r.Get("/value/{type}/{name}", handlers.GetMetric)
	r.Get("/", handlers.GetAllMetrics)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		panic(err)
	}

}
