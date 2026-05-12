package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Health check — first route every API needs
	// Used by Kubernetes liveness/readiness probes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "sameer-devops-platform-api",
		})
	})

	r.GET("/api/v1/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Sameer DevOps Platform API",
			"version": "0.1.0",
			"env":     os.Getenv("ENV"),
		})
	})

	log.Printf("API server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
