package prom

/*
  https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang
*/
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewPromMiddleware() gin.HandlerFunc {
	totalRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)

	responseStatus := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)

	httpDuration := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_time_seconds",
			Help: "Duration of HTTP requests.",
		},
		[]string{"path"},
	)

	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)

	return func(ctx *gin.Context) {
		p := ctx.Request.URL.Path
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(p))

		ctx.Next()

		responseStatus.WithLabelValues(strconv.Itoa(ctx.Writer.Status())).Inc()
		totalRequests.WithLabelValues(p).Inc()

		timer.ObserveDuration()
	}
}

func NewPromHandler() gin.HandlerFunc {
	return WrapH(promhttp.Handler())
}

func WrapF(f http.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f(ctx.Writer, ctx.Request)
	}
}

func WrapH(h http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
