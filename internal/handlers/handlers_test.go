package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMetric(t *testing.T) {
	tests := []struct {
		testName string
		urlType  string
		urlName  string
		urlValue string
		want     int
	}{
		{
			testName: "simple test with gauge type",
			urlType:  `gauge`,
			urlName:  `testGauge`,
			urlValue: `-12.34`,
			want:     http.StatusOK,
		},
		{
			testName: "bad type",
			urlType:  `badType`,
			urlName:  `test`,
			urlValue: `100`,
			want:     http.StatusBadRequest,
		},
		{
			testName: "bad value of counter",
			urlType:  `counter`,
			urlName:  `testCounter`,
			urlValue: `5ive`,
			want:     http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/update/{type}/{name}/{value}", nil)
			req.SetPathValue("type", test.urlType)
			req.SetPathValue("name", test.urlName)
			req.SetPathValue("value", test.urlValue)
			w := httptest.NewRecorder()
			AddMetric(w, req)
			res := w.Result()
			assert.Equal(t, test.want, res.StatusCode)
			res.Body.Close()
		},
		)

	}

}
