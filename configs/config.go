package configs

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort          string
	DBHost           string
	DBPort           int
	DBUser           string
	DBPassword       string
	DBName           string
	DBSSLMode        string
	DBTimeZone       string
	DBMaxOpenConns   int
	DBMaxIdleConns   int
	DBConnMaxLifetime time.Duration
}

func getenv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func atoi(key string, def int) int {
	v := getenv(key, "")
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}

func dur(key string, def time.Duration) time.Duration {
	v := getenv(key, "")
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return def
	}
	return d
}

func Load() *Config {
	return &Config{
		AppPort:           getenv("APP_PORT", "8080"),
		DBHost:            getenv("DB_HOST", "localhost"),
		DBPort:            atoi("DB_PORT", 5432),
		DBUser:            getenv("DB_USER", "dev"),
		DBPassword:        getenv("DB_PASSWORD", "devpass"),
		DBName:            getenv("DB_NAME", "appdb"),
		DBSSLMode:         getenv("DB_SSLMODE", "disable"),
		DBTimeZone:        getenv("DB_TIMEZONE", "America/Bogota"),
		DBMaxOpenConns:    atoi("DB_MAX_OPEN_CONNS", 20),
		DBMaxIdleConns:    atoi("DB_MAX_IDLE_CONNS", 10),
		DBConnMaxLifetime: dur("DB_CONN_MAX_LIFETIME", 30*time.Minute),
	}
}
