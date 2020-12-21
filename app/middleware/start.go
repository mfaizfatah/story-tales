package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//GetStatusCode is struct
type GetStatusCode struct {
	http.ResponseWriter
	status int
}

//WriteHeader is func test to superflous
func (w *GetStatusCode) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// RecordMiddleware is using for record any request from outside
func RecordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Whitelist(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}

		start := time.Now().In(loc)

		req := logger.StartRecord(r, start)

		// sc := GetStatusCode{ResponseWriter: w}
		next.ServeHTTP(w, req)
		// EndProcessing(req, sc.status, start)
	})
}

/*EndProcessing parts
 * @updated: Tuesday, December 3rd, 2019.
 * --
 * @param	p    	*Prometheus
 * @param	start	time.Duration
 * @param	wr   	http.ResponseWriter
 * @param	r    	*http.Request
 * @return	void
 */
func EndProcessing(r *http.Request, status int, start time.Time) {
	var st string
	t := time.Since(start)

	st = strconv.Itoa(status)
	if status == 0 {
		st = "200"
	}

	utils.Getprometheus().MetricRecord(st, r.Method, r.URL.Path, "backend", t)
}

/*Whitelist parts
 * @updated: Tuesday, December 3rd, 2019.
 * --
 * @param	urls	string
 * @return	mixed
 */
func Whitelist(urls string) bool {
	ws := regexp.MustCompile(`^\/(health|metrics|favicon.ico|liveness|debug[\w|\W]*)`)
	return ws.MatchString(urls)
}

var (
	HttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "myapp_http_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: utils.DefaultBuckets,
	}, []string{"code", "method", "path"})
)

// PrometheusMiddleware implements mux.MiddlewareFunc.
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Whitelist(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		sc := GetStatusCode{ResponseWriter: w}
		code := strconv.Itoa(sc.status)
		timer := prometheus.NewTimer(HttpDuration.WithLabelValues(code, r.URL.Path, r.Method))
		next.ServeHTTP(&sc, r)
		timer.ObserveDuration()
	})
}
