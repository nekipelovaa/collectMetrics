package storage

import "fmt"

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func (m *MemStorage) AddGaugeMetric(name string, val float64) {
	m.gauge[name] = val
}

func (m *MemStorage) AddCounterMetric(name string, val int64) {
	m.counter[name] += val
}

func NewStorage() *MemStorage {
	return &MemStorage{gauge: make(map[string]float64), counter: make(map[string]int64)}
}

func (m *MemStorage) GetCounterMetric(name string) (int64, bool) {
	v, ok := m.counter[name]
	return v, ok
}

func (m *MemStorage) GetGaugeMetric(name string) (float64, bool) {
	v, ok := m.gauge[name]
	return v, ok
}

func (m *MemStorage) GetAllMetricsToStr() string {
	out := "Gauge metrics:\r\n"
	for n, v := range m.gauge {
		out += fmt.Sprintf("\t%s\t=\t%f\r\n", n, v)
	}
	out += "Counter metrics:\r\n"
	for n, v := range m.counter {
		out += fmt.Sprintf("\t%s\t=\t%d\r\n", n, v)
	}
	return out
}
