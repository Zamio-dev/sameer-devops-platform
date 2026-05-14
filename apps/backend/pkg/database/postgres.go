package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver — blank import registers it
	"go.uber.org/zap"

	"github.com/sameermalik/devops-platform/pkg/logger"
)

func NewPostgres(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Connection pool settings — critical for production performance
	// Without these, Go opens unlimited connections and crashes your DB
	db.SetMaxOpenConns(25)       // max simultaneous connections
	db.SetMaxIdleConns(10)       // keep 10 connections warm in the pool
	db.SetConnMaxLifetime(5 * time.Minute) // recycle connections every 5min

	// Verify the connection works
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres ping failed: %w", err)
	}

	logger.Log.Info("PostgreSQL connected",
		zap.String("host", "configured via DSN"),
		zap.Int("max_connections", 25),
	)

	return db, nil
}
