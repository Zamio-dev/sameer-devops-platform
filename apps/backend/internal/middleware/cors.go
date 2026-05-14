package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS allows the frontend (localhost:3000) to call the API (localhost:8080).
// Browsers block cross-origin requests by default — this unlocks it safely.
func CORS(allowedOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", allowedOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")

		// Preflight requests — browser sends OPTIONS before the real request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
