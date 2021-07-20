package prometheus_midleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "apilog_http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "apilog_response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "apilog_http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "apilog_cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

//func recordMetrics() {
//    go func() {
//        for {
//            opsProcessed.Inc()
//            time.Sleep(2 * time.Second)
//        }
//    }()
//}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "apilog_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		cpuTemp.Set(65.3)
		statusCode := c.Writer.Status()
		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	}
}
