package handlers

import (
	"net/http"
	"strconv"

	"github.com/nekipelovaa/collectMetrics/internal/storage"
)

var m *storage.MemStorage = storage.NewStorage()

func AddMetric(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if name == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.PathValue("type") {
	case "gauge":
		val, err := strconv.ParseFloat(r.PathValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		m.AddGaugeMetric(name, val)
	case "counter":
		val, err := strconv.ParseInt(r.PathValue("value"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		m.AddCounterMetric(name, val)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// var AddMetric func () http.Handler = http.HandlerFunc(func AddMetric(w http.ResponseWriter, r *http.Request) {
// 	name := r.PathValue("name")

// 	if name == "" {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	switch r.PathValue("type") {
// 	case "gauge":
// 		val, err := strconv.ParseFloat(r.PathValue("value"), 64)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		m.addGaugeMetric(name, val)
// 	case "counter":
// 		val, err := strconv.ParseInt(r.PathValue("value"), 10, 64)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		m.addCounterMetric(name, val)
// 	default:
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// })
