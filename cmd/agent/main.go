package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type MetricsCollection struct {
	gougeMetrics   map[string]float64
	counterMetrics map[string]int64
	sync.Mutex
}

func MetricsCollectionInit() *MetricsCollection {
	return &MetricsCollection{
		gougeMetrics:   make(map[string]float64),
		counterMetrics: make(map[string]int64),
	}
}

func (c *MetricsCollection) CollectMetrics() {
	c.Lock()
	mStats := &runtime.MemStats{}
	runtime.ReadMemStats(mStats)
	c.gougeMetrics["Alloc"] = float64(mStats.Alloc)
	c.gougeMetrics["BuckHashSys"] = float64(mStats.BuckHashSys)
	c.gougeMetrics["Frees"] = float64(mStats.Frees)
	c.gougeMetrics["GCCPUFraction"] = float64(mStats.GCCPUFraction)
	c.gougeMetrics["GCSys"] = float64(mStats.GCSys)
	c.gougeMetrics["HeapAlloc"] = float64(mStats.HeapAlloc)
	c.gougeMetrics["HeapIdle"] = float64(mStats.HeapIdle)
	c.gougeMetrics["HeapInuse"] = float64(mStats.HeapInuse)
	c.gougeMetrics["HeapObjects"] = float64(mStats.HeapObjects)
	c.gougeMetrics["HeapReleased"] = float64(mStats.HeapReleased)
	c.gougeMetrics["HeapSys"] = float64(mStats.HeapSys)
	c.gougeMetrics["LastGC"] = float64(mStats.LastGC)
	c.gougeMetrics["Lookups"] = float64(mStats.Lookups)
	c.gougeMetrics["MCacheInuse"] = float64(mStats.MCacheInuse)
	c.gougeMetrics["MCacheSys"] = float64(mStats.MCacheSys)
	c.gougeMetrics["MSpanInuse"] = float64(mStats.MSpanInuse)
	c.gougeMetrics["MSpanSys"] = float64(mStats.MSpanSys)
	c.gougeMetrics["Mallocs"] = float64(mStats.Mallocs)
	c.gougeMetrics["NextGC"] = float64(mStats.NextGC)
	c.gougeMetrics["NumForcedGC"] = float64(mStats.NumForcedGC)
	c.gougeMetrics["NumGC"] = float64(mStats.NumGC)
	c.gougeMetrics["OtherSys"] = float64(mStats.OtherSys)
	c.gougeMetrics["PauseTotalNs"] = float64(mStats.PauseTotalNs)
	c.gougeMetrics["StackInuse"] = float64(mStats.StackInuse)
	c.gougeMetrics["StackSys"] = float64(mStats.StackSys)
	c.gougeMetrics["Sys"] = float64(mStats.Sys)
	c.gougeMetrics["TotalAlloc"] = float64(mStats.TotalAlloc)

	c.gougeMetrics["RandomValue"] = rand.Float64() * 100
	c.counterMetrics["PollCount"] += 1
	//fmt.Printf("%v", c.counterMetrics["PollCount"])
	c.Unlock()
}

func main() {
	pollInterval := 2
	reportInterval := 10
	m := MetricsCollectionInit()

	go func() {
		for {
			time.Sleep(time.Duration(pollInterval) * time.Second)
			m.CollectMetrics()

		}
	}()
	for {
		time.Sleep(time.Duration(reportInterval) * time.Second)
		m.Lock()
		for k, v := range m.gougeMetrics {
			url := fmt.Sprintf("http://localhost:8080/update/gauge/%s/%v", k, v)
			resp, err := http.Post(url, "text/plain", nil)
			if err != nil {
				fmt.Println(err)
			}
			resp.Body.Close()
		}
		for k, v := range m.counterMetrics {
			url := fmt.Sprintf("http://localhost:8080/update/counter/%s/%v", k, v)
			resp, err := http.Post(url, "text/plain", nil)
			if err != nil {
				fmt.Println(err)
			}
			resp.Body.Close()
		}
		m.Unlock()
	}
}
