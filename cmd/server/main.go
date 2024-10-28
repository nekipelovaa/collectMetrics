package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/nekipelovaa/collectMetrics/internal/handlers"
	lh "github.com/nekipelovaa/collectMetrics/internal/loggingHandlers"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "localhost:8080", "адрес HTTP сервера")
	flag.Parse()

	addrEnv := os.Getenv("ADDRESS")
	if addrEnv != "" {
		addr = addrEnv
	}

	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", lh.WithLogging(handlers.AddMetric))
	r.Get("/value/{type}/{name}", lh.WithLogging(handlers.GetMetric))
	r.Get("/", lh.WithLogging(handlers.GetAllMetrics))
	err := http.ListenAndServe(addr, r)
	fmt.Println(addr)
	if err != nil {
		panic(err)
	}

}
