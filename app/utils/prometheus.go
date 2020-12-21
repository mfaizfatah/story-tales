package utils

import (
	"time"

	_ "github.com/joho/godotenv/autoload" //buat jaga2

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// DefaultBuckets prometheus buckets in seconds.
	DefaultBuckets = []float64{0.3, 1.2, 5.0}
)

const (
	reqsName    = "http_requests_total"
	latencyName = "http_request_duration_seconds"
)

// Prometheus is a handler that exposes prometheus metrics for the number of requests,
// the latency and the response size, partitioned by status code, method and HTTP path.
//
// Usage: pass its `ServeHTTP` to a route or globally.
type Prometheus struct {
	reqs    *prometheus.CounterVec
	latency *prometheus.HistogramVec
}

var ptheus *Prometheus

/*Newprometheus parts
 * @updated: Monday, December 2nd, 2019.
 * --
 * @param	name   	string
 * @param	buckets	...float64
 * @return	mixed
 */
func Newprometheus(name string) {
	str := []string{"code", "method", "path", "type"}
	p := Prometheus{}
	p.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        reqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"service": name},
		},
		str,
	)
	// prometheus.MustRegister(p.reqs)

	p.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        latencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
		ConstLabels: prometheus.Labels{"service": name},
		Buckets:     DefaultBuckets,
	},
		str,
	)
	// prometheus.MustRegister(p.latency)
	ptheus = &p
}

//MetricRecord is recorder
func (p Prometheus) MetricRecord(code, method, path, types string, t time.Duration) {
	p.reqs.WithLabelValues(code, method, path, types).Inc()
	p.latency.WithLabelValues(code, method, path, types).Observe(float64(t.Nanoseconds()) / 1000000000)
}

/*Getprometheus parts
 * @updated: Wednesday, December 4th, 2019.
 * --
 * @return	void
 */
func Getprometheus() Prometheus {
	return *ptheus
}
