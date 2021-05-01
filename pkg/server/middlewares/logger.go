package middlewares

import (
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
)

func HTTPLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(h, rw, r)

		log.Printf(
			"--> %s - %s %s\n",
			r.Method,
			r.Host,
			r.URL.Path,
		)
		log.Printf(
			"<-- %d - %s %d\n",
			m.Code,
			m.Duration,
			m.Written,
		)
	})
}
