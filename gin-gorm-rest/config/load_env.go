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

	RefreshTokenExpiresIn time.Duration
	RefreshTokenMaxAge    int
	RefreshTokenSecret    string

	AccessTokenExpiresIn time.Duration
	AccessTokenSecret    string

	RedisHost string
	RedisPort string
	RedisDB   int
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
		"REFRESH_TOKEN_EXPIRATION", "REFRESH_TOKEN_MAXAGE", "REFRESH_TOKEN_SECRET",
		"ACCESS_TOKEN_EXPIRATION", "ACCESS_TOKEN_SECRET",
		"REDIS_HOST", "REDIS_PORT", "REDIS_DB",
	}
	for _, env := range requiredEnvVars {
		if os.Getenv(env) == "" {
			return Config{}, fmt.Errorf("Environment variable %s is not set", env)
		}
	}

	// Parse TOKEN_EXPIRATION (e.g., "60m", "2h", ...)
	refreshTokenExpiration, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRATION"))
	if err != nil {
		return Config{}, fmt.Errorf("Invalid format for REFRESH TOKEN_EXPIRATION: %v", err)
	}

	accessTokenExpiration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRATION"))
	if err != nil {
		return Config{}, fmt.Errorf("Invalid format for ACCESS TOKEN_EXPIRATION: %v", err)
	}

	// Parse TOKEN_MAXAGE
	refreshTokenMaxAge, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAXAGE"))
	if err != nil {
		return Config{}, fmt.Errorf("Invalid value for REFRESH_TOKEN_MAXAGE: %v", err)
	}

	// Parse REDIS_DB
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return Config{}, fmt.Errorf("Invalid value for REDIS_DB: %v", err)
	}

	// Return configuration struct
	return Config{
		PostgresUser:          os.Getenv("POSTGRES_USER"),
		PostgresPassword:      os.Getenv("POSTGRES_PASSWORD"),
		DBHost:                os.Getenv("DB_HOST"),
		DBPort:                os.Getenv("DB_PORT"),
		PostgresDB:            os.Getenv("POSTGRES_DB"),
		RefreshTokenExpiresIn: refreshTokenExpiration,
		RefreshTokenMaxAge:    refreshTokenMaxAge,
		RefreshTokenSecret:    os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenExpiresIn:  accessTokenExpiration,
		AccessTokenSecret:     os.Getenv("ACCESS_TOKEN_SECRET"),
		RedisHost:             os.Getenv("REDIS_HOST"),
		RedisPort:             os.Getenv("REDIS_PORT"),
		RedisDB:               redisDB,
	}, nil
}
