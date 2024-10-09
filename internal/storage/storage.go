package storage

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

// var m MemStorage = MemStorage{
// 	gauge:   make(map[string]float64),
// 	counter: make(map[string]int64),
// }

func NewStorage() *MemStorage {
	return &MemStorage{gauge: make(map[string]float64), counter: make(map[string]int64)}
}
