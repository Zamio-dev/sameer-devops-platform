package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type HealthHandler struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewHealthHandler(db *sqlx.DB, redis *redis.Client) *HealthHandler {
	return &HealthHandler{db: db, redis: redis}
}

type healthResponse struct {
	Status   string            `json:"status"`
	Service  string            `json:"service"`
	Checks   map[string]string `json:"checks"`
	Uptime   string            `json:"uptime"`
}

var startTime = time.Now()

// Live — liveness probe. Simple: is the process running?
func (h *HealthHandler) Live(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Ready — readiness probe. Checks all dependencies.
func (h *HealthHandler) Ready(c *gin.Context) {
	checks := map[string]string{}
	status := "ok"
	httpStatus := http.StatusOK

	// Check PostgreSQL
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := h.db.PingContext(ctx); err != nil {
		checks["postgres"] = "unhealthy: " + err.Error()
		status = "degraded"
		httpStatus = http.StatusServiceUnavailable
	} else {
		checks["postgres"] = "healthy"
	}

	// Check Redis
	if err := h.redis.Ping(ctx).Err(); err != nil {
		checks["redis"] = "unhealthy: " + err.Error()
		status = "degraded"
		httpStatus = http.StatusServiceUnavailable
	} else {
		checks["redis"] = "healthy"
	}

	c.JSON(httpStatus, healthResponse{
		Status:  status,
		Service: "sameer-devops-platform-api",
		Checks:  checks,
		Uptime:  time.Since(startTime).Round(time.Second).String(),
	})
}
