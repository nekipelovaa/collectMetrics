package main

import (
	"net/http"
	"strconv"
)

//	type Gayge struct {
//		float64
//	}
// type MemStorage struct {
// 	gaugeMetrics   map[string]float64
// 	counterMetrics map[string][]int64
// }

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func (m *MemStorage) addGaugeMetric(name string, val float64) {
	m.gauge[name] = val
}
func (m *MemStorage) addCounterMetric(name string, val int64) {
	m.counter[name] += val
}

var m MemStorage = MemStorage{
	gauge:   make(map[string]float64),
	counter: make(map[string]int64),
}

func addMetric(w http.ResponseWriter, r *http.Request) {
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
		m.addGaugeMetric(name, val)
	case "counter":
		val, err := strconv.ParseInt(r.PathValue("value"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		m.addCounterMetric(name, val)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /update/{type}/{name}/{value}/", addMetric)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
