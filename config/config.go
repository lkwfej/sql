package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds runtime configuration.
// Keep it simple for a demo project: read from env with sane defaults.
type Config struct {
	// Server
	Addr    string // e.g. ":9090"
	GinMode string // debug / release

	// Database
	MySQLDSN string

	// JWT
	JWTSecret []byte
	TokenTTL  time.Duration
}

func Load() Config {
	addr := getenv("ADDR", ":9090")
	ginMode := getenv("GIN_MODE", "debug")

	// Default DSN matches your current working setup.
	dsn := getenv("MYSQL_DSN", "admin_sql:admin_sql@tcp(127.0.0.1:3306)/goland_demo?charset=utf8mb4&parseTime=True&loc=Local")

	// !!! In production, ALWAYS set JWT_SECRET via env.
	secret := []byte(getenv("JWT_SECRET", "change_this_secret"))

	// Token TTL in hours (default: 24h)
	ttlHours := getenvInt("JWT_TTL_HOURS", 24)
	if ttlHours <= 0 {
		ttlHours = 24
	}

	return Config{
		Addr:      addr,
		GinMode:   ginMode,
		MySQLDSN:  dsn,
		JWTSecret: secret,
		TokenTTL:  time.Duration(ttlHours) * time.Hour,
	}
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func getenvInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
