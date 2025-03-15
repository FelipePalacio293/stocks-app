package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	ServerPort        string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	AllowedOrigins    []string
	Environment       string
	EnableRequestLogs bool
	APIBaseURL        string
	APIKey            string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	enableLogs, err := strconv.ParseBool(getEnv("ENABLE_REQUEST_LOGS", "true"))
	if err != nil {
		return nil, fmt.Errorf("invalid ENABLE_REQUEST_LOGS: %w", err)
	}

	return &Config{
		ServerPort:        getEnv("SERVER_PORT", "8080"),
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "26257"),
		DBUser:            getEnv("DB_USER", "root"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "go_api"),
		AllowedOrigins:    []string{getEnv("ALLOWED_ORIGINS", "*")},
		Environment:       getEnv("ENVIRONMENT", "development"),
		EnableRequestLogs: enableLogs,
		APIBaseURL:        getEnv("API_BASE_URL", ""),
		APIKey:            getEnv("API_KEY", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var logLevel logger.LogLevel
	if cfg.Environment == "production" {
		logLevel = logger.Error
	} else {
		logLevel = logger.Info
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	return gorm.Open(postgres.Open(dsn), config)
}
