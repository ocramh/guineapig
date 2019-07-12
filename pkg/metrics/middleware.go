package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsResponseTime prometheus.Summary
	httpRequestsStatus       *prometheus.CounterVec
)

//////
// SEE: https://github.com/slok/go-http-metrics
//////

func init() {
	httpRequestsResponseTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "http",
		Name:      "http_response_time_seconds",
		Help:      "Request response times",
	})

	httpRequestsStatus = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http",
		Name:      "http_requests_total",
		Help:      "How many HTTP requests processed, partitioned by status code and HTTP method",
	},
		[]string{"code", "method"})

	prometheus.MustRegister(httpRequestsResponseTime)
	prometheus.MustRegister(httpRequestsStatus)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, 200}
		start := time.Now()

		next.ServeHTTP(&rec, r)

		httpRequestsResponseTime.Observe(float64(time.Since(start).Seconds()))
		httpRequestsStatus.WithLabelValues(strconv.Itoa(rec.status), r.Method).Add(1)
	})
}
