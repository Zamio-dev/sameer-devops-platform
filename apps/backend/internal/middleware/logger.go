package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sameermalik/devops-platform/pkg/logger"
)

// Logger logs every HTTP request with method, path, status, latency.
// This is the access log — every production API has one.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logger.Log.Info("request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("request_id", c.GetString("request_id")),
			zap.String("ip", c.ClientIP()),
		)
	}
}
