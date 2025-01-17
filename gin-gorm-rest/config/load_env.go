package config

import (
	"fmt"

	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	DBHost           string
	DBPort           string
	PostgresDB       string

	TokenExpiresIn time.Duration
	TokenMaxAge    int
	TokenSecret    string
}

// LoadConfig tải các thông số cấu hình từ file .env vào struct Config
// LoadConfig loads configuration from .env file and returns it with an error if any.
func LoadConfig() (Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("Error loading .env file: %v", err)
	}

	// Check for required environment variables
	requiredEnvVars := []string{
		"POSTGRES_USER", "POSTGRES_PASSWORD", "DB_HOST", "DB_PORT", "POSTGRES_DB",
		"TOKEN_EXPIRATION", "TOKEN_MAXAGE", "TOKEN_SECRET",
	}
	for _, env := range requiredEnvVars {
		if os.Getenv(env) == "" {
			return Config{}, fmt.Errorf("Environment variable %s is not set", env)
		}
	}

	// Parse TOKEN_EXPIRATION (e.g., "60m", "2h", ...)
	tokenExpiration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION"))
	if err != nil {
		return Config{}, fmt.Errorf("Invalid format for TOKEN_EXPIRATION: %v", err)
	}

	// Parse TOKEN_MAXAGE
	tokenMaxAge, err := strconv.Atoi(os.Getenv("TOKEN_MAXAGE"))
	if err != nil {
		return Config{}, fmt.Errorf("Invalid value for TOKEN_MAXAGE: %v", err)
	}

	// Return configuration struct
	return Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBHost:           os.Getenv("DB_HOST"),
		DBPort:           os.Getenv("DB_PORT"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		TokenExpiresIn:   tokenExpiration,
		TokenMaxAge:      tokenMaxAge,
		TokenSecret:      os.Getenv("TOKEN_SECRET"),
	}, nil
}
