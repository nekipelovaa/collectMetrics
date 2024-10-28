package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/nekipelovaa/collectMetrics/internal/storage"
)

var m *storage.MemStorage = storage.NewStorage()

func AddMetric(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

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

func GetMetric(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	typeMetric := r.PathValue("type")

	switch typeMetric {
	case "counter":
		v, ok := m.GetCounterMetric(name)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain")
		_, err := w.Write([]byte(fmt.Sprintf("%d", v)))
		if err != nil {
			fmt.Println(err)
		}
		return
	case "gauge":
		v, ok := m.GetGaugeMetric(name)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain")
		_, err := w.Write([]byte(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", v), "0"), ".")))
		if err != nil {
			fmt.Println(err)
		}
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func GetAllMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, err := w.Write([]byte(m.GetAllMetricsToStr()))
	if err != nil {
		fmt.Println(err)
	}
}
