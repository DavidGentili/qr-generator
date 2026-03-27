package observability

import (
	"log"
	"net/http"
	"qr-generator/config"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requests HTTP procesadas por la API",
		},
		[]string{"method", "path", "status"},
	)
	httpRequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duración de requests HTTP en segundos",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)
	qrGeneratedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "qr_generated_total",
			Help: "Total de códigos QR generados correctamente",
		},
	)
	qrGenerationErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "qr_generation_errors_total",
			Help: "Total de errores durante el flujo de generación QR",
		},
		[]string{"stage"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDurationSeconds, qrGeneratedTotal, qrGenerationErrorsTotal)
}

func HTTPMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}
		status := strconv.Itoa(c.Writer.Status())

		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
		httpRequestDurationSeconds.WithLabelValues(method, path, status).Observe(time.Since(start).Seconds())
	}
}

func StartMetricsServer(observabilityParams config.ObservabilityParams) {
	if !observabilityParams.MetricsEnabled {
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Printf("[Observability] Iniciando servidor de métricas en %s", observabilityParams.MetricsPort)
		if err := http.ListenAndServe(observabilityParams.MetricsPort, mux); err != nil {
			log.Printf("[Observability] Error en servidor de métricas: %v", err)
		}
	}()
}

func IncQRGenerated() {
	qrGeneratedTotal.Inc()
}

func IncQRGenerationError(stage string) {
	qrGenerationErrorsTotal.WithLabelValues(stage).Inc()
}
