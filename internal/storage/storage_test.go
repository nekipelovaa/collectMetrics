package storage

import (
	"testing"
)

func TestNewStorage(t *testing.T) {
	m := NewStorage()
	if m == nil {
		t.Errorf("Ожидается, что NewStorage вернет ненулевой экземпляр MemStorage")
	}
	if len(m.gauge) != 0 {
		t.Errorf("Ожидалось, что мапа Gauge будет пустой, получено %v", m.gauge)
	}
	if len(m.counter) != 0 {
		t.Errorf("Ожидалось, что мапа Counter будет пустой, получено %v", m.counter)
	}
}

func TestAddGaugeMetric(t *testing.T) {
	m := NewStorage()
	m.AddGaugeMetric("test", -10.5)

	if _, ok := m.gauge["test"]; !ok {
		t.Errorf("Gauge метрика  не была сохранена")
	}
	if m.gauge["test"] != -10.5 {
		t.Errorf("Ожидание значения метрики -10.5, получен %f", m.gauge["test"])
	}
	m.AddGaugeMetric("test", 8.75)
	if m.gauge["test"] != 8.75 {
		t.Errorf("После повторного добавления метрики, ожидание значения 8.75, получили %f", m.gauge["test"])
	}
}

func TestAddCounterMetric(t *testing.T) {
	m := NewStorage()
	m.AddCounterMetric("test", 10)

	if _, ok := m.counter["test"]; !ok {
		t.Errorf("Counter метрика  не была сохранена")
	}
	if m.counter["test"] != 10 {
		t.Errorf("Ожидание значения метрики 10, получен  %d", m.counter["test"])
	}

	m.AddCounterMetric("test", 20)
	if m.counter["test"] != 30 {
		t.Errorf("После повторного добавления метрики, ожидание значения 30, получили %d", m.counter["test"])
	}
}
