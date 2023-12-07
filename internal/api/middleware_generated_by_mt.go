// Code generated by mtgroup-generator.
package rest

import (
	"net"
	"net/http"
	"strings"

	"github.com/felixge/httpsnoop"
	"github.com/powerman/structlog"
	"go.uber.org/zap"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

type middlewareFunc func(http.Handler) http.Handler

// Provide a logger configured using request's context.
//
// Usually it should be first middleware.
func makeLogger(basePath string, l *zap.SugaredLogger) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Infow("<-", "remote_addr", r.RemoteAddr, "method", r.Method, "func", strings.TrimPrefix(r.URL.Path, basePath))
			next.ServeHTTP(w, r)
		})
	}
}

// go-swagger responders panic on error while writing response to client,
// this shouldn't result in crash - unlike a real, reasonable panic.
//
// Usually it should be second middleware (after logger).
func recovery(next http.Handler, l *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			const code = http.StatusInternalServerError
			switch err := recover(); err := err.(type) {
			default:
				l.Errorw("panic", "http_status", code, "err", err, structlog.KeyStack, structlog.Auto)
				w.WriteHeader(code)
			case nil:
			case net.Error:
				l.Errorw("panic", "http_status", code, "err", err)
				w.WriteHeader(code)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func makeAccessLog(basePath string, l *zap.SugaredLogger) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m := httpsnoop.CaptureMetrics(next, w, r)

			if m.Code < 500 {
				l.Infow("handled", "http_status", m.Code)
			} else {
				l.Infow("failed to handle", "http_status", m.Code)
			}
		})
	}
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Expires", "0")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		next.ServeHTTP(w, r)
	})
}