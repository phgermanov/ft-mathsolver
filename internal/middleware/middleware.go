package middleware

import (
	"net/http"
	"time"

	"github.com/phgermanov/ft-mathsolver/internal"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		internal.Log.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL,
			"client": r.RemoteAddr,
		}).Info("Started request")

		next.ServeHTTP(w, r)

		internal.Log.WithFields(logrus.Fields{
			"method":   r.Method,
			"url":      r.URL,
			"client":   r.RemoteAddr,
			"duration": time.Since(start),
		}).Info("Finished request")
	})
}
