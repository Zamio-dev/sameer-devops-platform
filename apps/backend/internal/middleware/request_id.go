package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID attaches a unique ID to every request.
// This ID flows through all logs for that request — essential for debugging.
// "Why did this request fail?" → grep by request ID → see every log line.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Set("request_id", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}
