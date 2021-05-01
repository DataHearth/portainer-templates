package middlewares

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Prometheus interface {
	IncreaseRequest(statusCode string)
	SaveDurationHTTP(method, path, statusCode, dataWritten string, duration float64)
}

//Service implements UseCase interface
type prom struct {
	logger               logrus.FieldLogger
	durationHistogramVec *prometheus.HistogramVec
	requestCounterVec    *prometheus.CounterVec
}

//NewPrometheusService create a new prometheus service
func NewPrometheusService(logger logrus.FieldLogger) (Prometheus, error) {
	duration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "Latency of a HTTP request with [method, path, code, data_written]",
		Buckets:   prometheus.DefBuckets,
	}, []string{"method", "path", "code", "data_written"})
	request := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http",
		Name:      "request_counter",
		Help:      "Numbers of request",
	}, []string{"code"})

	p := &prom{
		logger:               logger.WithField("pkg", "middlewares"),
		durationHistogramVec: duration,
		requestCounterVec:    request,
	}

	if err := prometheus.Register(p.durationHistogramVec); err != nil {
		return nil, err
	}
	if err := prometheus.Register(p.requestCounterVec); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *prom) IncreaseRequest(statusCode string) {
	counter, err := p.requestCounterVec.GetMetricWithLabelValues(statusCode)
	if err != nil {
		p.logger.WithField("component", "IncreaseRequest").Errorln("failed to get request counter metric")
	}

	counter.Inc()
}

func (p *prom) SaveDurationHTTP(method, path, statusCode, dataWritten string, duration float64) {
	histo, err := p.durationHistogramVec.GetMetricWithLabelValues(method, path, statusCode, dataWritten)
	if err != nil {
		p.logger.WithError(err).WithField("component", "SaveDurationHTTP").Errorln("failed to get duration histogram metric")
		return
	}

	histo.Observe(duration)
}
