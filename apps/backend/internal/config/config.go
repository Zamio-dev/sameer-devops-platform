package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Port string
	Env  string
	Name string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
	DSN      string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       string
}

func Load() *Config {
	// Load .env file if it exists (dev only — production uses real env vars)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	cfg := &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8080"),
			Env:  getEnv("ENV", "development"),
			Name: getEnv("APP_NAME", "sameer-devops-platform"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "platform_db"),
			User:     getEnv("DB_USER", "platform_user"),
			Password: getEnv("DB_PASSWORD", "platform_dev_password"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
	}

	// Build DSN (Data Source Name) — the PostgreSQL connection string
	cfg.Database.DSN = "host=" + cfg.Database.Host +
		" port=" + cfg.Database.Port +
		" dbname=" + cfg.Database.Name +
		" user=" + cfg.Database.User +
		" password=" + cfg.Database.Password +
		" sslmode=" + cfg.Database.SSLMode

	return cfg
}

// getEnv reads env var, returns fallback if not set
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
