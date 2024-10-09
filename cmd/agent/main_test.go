package main

import (
	"testing"
)

func TestMetricsCollectionInit(t *testing.T) {
	m := MetricsCollectionInit()
	if m == nil {
		t.Errorf("Ожидалось что функция вернет не nil")
	}
}

func TestCollectMetrics(t *testing.T) {
	m := MetricsCollectionInit()
	m.CollectMetrics()

	if len(m.gougeMetrics) == 0 {
		t.Errorf("Ожидается не нулевая мапа gaugeMetrics")
	}
	if len(m.counterMetrics) == 0 {
		t.Errorf("Ожидается не нулевая мапа counterMetrics")
	}

	if _, ok := m.gougeMetrics["RandomValue"]; !ok {
		t.Errorf("Ожидается значение в метрике RandomValue")
	}

	if _, ok := m.counterMetrics["PollCount"]; !ok {
		t.Errorf("Ожидается значение в метрике PollCount")
	}

	m.CollectMetrics()
	if m.counterMetrics["PollCount"] != 2 {
		t.Errorf("Ожидается в PullCounter значение 2, получено %d", m.counterMetrics["PollCount"])
	}
}
