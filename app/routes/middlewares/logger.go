package middlewares

import (
	"bytes"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	const funcName string = "responseWriter.Write"

	if _, err := rw.body.Write(b); err != nil {
		slog.Error("write error",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		return 0, err
	}

	return rw.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	const funcName string = "LoggerMiddleware"

	return func(c *gin.Context) {
		start := time.Now()

		slog.Info("HTTP Request",
			slog.String("function_name", funcName),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("query", c.Request.URL.RawQuery),
		)

		rw := &responseWriter{ResponseWriter: c.Writer,
			body: bytes.NewBufferString("")}
		c.Writer = rw

		c.Next()

		duration := time.Since(start)

		slog.Info("HTTP Response",
			slog.String("function_name", funcName),
			slog.Int("status", rw.Status()),
			slog.String("duration", duration.String()),
		)
	}
}
