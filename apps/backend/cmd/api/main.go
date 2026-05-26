package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sameermalik/devops-platform/internal/config"
	"github.com/sameermalik/devops-platform/internal/handler"
	"github.com/sameermalik/devops-platform/internal/middleware"
	"github.com/sameermalik/devops-platform/pkg/database"
	"github.com/sameermalik/devops-platform/pkg/logger"
)

func main() {
	// 1. Load config
	cfg := config.Load()

	// 2. Init logger first — everything else uses it
	logger.Init(cfg.App.Env)
	defer logger.Sync()

	logger.Log.Info("starting server",
		zap.String("service", cfg.App.Name),
		zap.String("env", cfg.App.Env),
		zap.String("port", cfg.App.Port),
	)

	// 3. Connect to PostgreSQL
	db, err := database.NewPostgres(cfg.Database.DSN)
	if err != nil {
		logger.Log.Fatal("postgres connection failed", zap.Error(err))
	}
	defer db.Close()

	// Run database migrations
	if err := database.RunMigrations(db, "./migrations"); err != nil {
		logger.Log.Fatal("migrations failed", zap.Error(err))
	}

	// 4. Connect to Redis
	rdb, err := database.NewRedis(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		logger.Log.Fatal("redis connection failed", zap.Error(err))
	}
	defer rdb.Close()

	// 5. Set Gin mode
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 6. Build router
	r := gin.New() // gin.New() = no default middleware (we add our own)

	// Global middleware — order matters
	r.Use(middleware.RequestID()) // first: assign ID to request
	r.Use(middleware.Logger())    // second: log with that ID
	r.Use(middleware.CORS("http://localhost:3000"))
	r.Use(gin.Recovery()) // last: catch panics, return 500

	// 7. Register routes
	healthHandler := handler.NewHealthHandler(db, rdb)

	r.GET("/health", healthHandler.Live)
	r.GET("/ready",  healthHandler.Ready)

	api := r.Group("/api/v1")
	{
		api.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name":    cfg.App.Name,
				"version": "0.1.0",
				"env":     cfg.App.Env,
			})
		})
	}

	// 8. Start server with graceful shutdown
	// Why graceful shutdown? In-flight requests complete before server stops.
	// Without it, Kubernetes rolling updates drop live requests.
	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start in goroutine so we can listen for shutdown signal
	go func() {
		logger.Log.Info("server listening", zap.String("port", cfg.App.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("server error", zap.Error(err))
		}
	}()

	// Block until OS sends SIGINT (Ctrl+C) or SIGTERM (Kubernetes shutdown)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("forced shutdown", zap.Error(err))
	}

	logger.Log.Info("server stopped")
}
